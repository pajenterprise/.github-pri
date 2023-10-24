// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

package forwarder

import (
	"bytes"
	"context"
	"crypto/tls"
	"expvar"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"strconv"
	"time"

	"github.com/DataDog/datadog-agent/pkg/config"
	"github.com/DataDog/datadog-agent/pkg/telemetry"
	httputils "github.com/DataDog/datadog-agent/pkg/util/http"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

var (
	connectionDNSSuccess               = expvar.Int{}
	connectionConnectSuccess           = expvar.Int{}
	transactionsExpvars                = expvar.Map{}
	transactionsInputBytesByEndpoint   = expvar.Map{}
	transactionsConnectionEvents       = expvar.Map{}
	transactionsInputCountByEndpoint   = expvar.Map{}
	transactionsDropped                = expvar.Int{}
	transactionsDroppedByEndpoint      = expvar.Map{}
	transactionsDroppedOnInput         = expvar.Int{}
	transactionsRequeued               = expvar.Int{}
	transactionsRequeuedByEndpoint     = expvar.Map{}
	transactionsRetried                = expvar.Int{}
	transactionsRetriedByEndpoint      = expvar.Map{}
	transactionsRetryQueueSize         = expvar.Int{}
	transactionsSuccessByEndpoint      = expvar.Map{}
	transactionsSuccessBytesByEndpoint = expvar.Map{}
	transactionsSuccess                = expvar.Int{}
	transactionsErrors                 = expvar.Int{}
	transactionsErrorsByType           = expvar.Map{}
	transactionsDNSErrors              = expvar.Int{}
	transactionsTLSErrors              = expvar.Int{}
	transactionsConnectionErrors       = expvar.Int{}
	transactionsWroteRequestErrors     = expvar.Int{}
	transactionsSentRequestErrors      = expvar.Int{}
	transactionsHTTPErrors             = expvar.Int{}
	transactionsHTTPErrorsByCode       = expvar.Map{}

	tlmTxInputBytes = telemetry.NewCounter("transactions", "input_bytes",
		[]string{"domain", "endpoint"}, "Incoming transaction sizes in bytes")
	tlmConnectEvents = telemetry.NewCounter("transactions", "connection_events",
		[]string{"connection_event_type"}, "Count of new connection events grouped by type of event")
	tlmTxInputCount = telemetry.NewCounter("transactions", "input_count",
		[]string{"domain", "endpoint"}, "Incoming transaction count")
	tlmTxDropped = telemetry.NewCounter("transactions", "dropped",
		[]string{"domain", "endpoint"}, "Transaction drop count")
	tlmTxDroppedOnInput = telemetry.NewCounter("transactions", "dropped_on_input",
		[]string{"domain", "endpoint"}, "Count of transactions dropped on input")
	tlmTxRequeued = telemetry.NewCounter("transactions", "requeued",
		[]string{"domain", "endpoint"}, "Transaction requeue count")
	tlmTxRetried = telemetry.NewCounter("transactions", "retries",
		[]string{"domain", "endpoint"}, "Transaction retry count")
	tlmTxRetryQueueSize = telemetry.NewGauge("transactions", "retry_queue_size",
		[]string{"domain"}, "Retry queue size")
	tlmTxSuccessCount = telemetry.NewCounter("transactions", "success",
		[]string{"domain", "endpoint"}, "Successful transaction count")
	tlmTxSuccessBytes = telemetry.NewCounter("transactions", "success_bytes",
		[]string{"domain", "endpoint"}, "Successful transaction sizes in bytes")
	tlmTxErrors = telemetry.NewCounter("transactions", "errors",
		[]string{"domain", "endpoint", "error_type"}, "Count of transactions errored grouped by type of error")
	tlmTxHTTPErrors = telemetry.NewCounter("transactions", "http_errors",
		[]string{"domain", "endpoint", "code"}, "Count of transactions http errors per http code")
)

var trace = &httptrace.ClientTrace{
	DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
		if dnsInfo.Err != nil {
			transactionsDNSErrors.Add(1)
			tlmTxErrors.Inc("unknown", "unknown", "dns_lookup_failure")
			log.Debugf("DNS Lookup failure: %s", dnsInfo.Err)
			return
		}
		connectionDNSSuccess.Add(1)
		tlmConnectEvents.Inc("dns_lookup_success")
		log.Tracef("DNS Lookup success, addresses: %s", dnsInfo.Addrs)
	},
	WroteRequest: func(wroteInfo httptrace.WroteRequestInfo) {
		if wroteInfo.Err != nil {
			transactionsWroteRequestErrors.Add(1)
			tlmTxErrors.Inc("unknown", "unknown", "writing_failure")
			log.Debugf("Request writing failure: %s", wroteInfo.Err)
		}
	},
	ConnectDone: func(network, addr string, err error) {
		if err != nil {
			transactionsConnectionErrors.Add(1)
			tlmTxErrors.Inc("unknown", "unknown", "connection_failure")
			log.Debugf("Connection failure: %s", err)
			return
		}
		connectionConnectSuccess.Add(1)
		tlmConnectEvents.Inc("connection_success")
		log.Tracef("New successful connection to address: %q", addr)
	},
	TLSHandshakeDone: func(tlsState tls.ConnectionState, err error) {
		if err != nil {
			transactionsTLSErrors.Add(1)
			tlmTxErrors.Inc("unknown", "unknown", "tls_handshake_failure")
			log.Errorf("TLS Handshake failure: %s", err)
		}
	},
}

// Compile-time check to ensure that HTTPTransaction conforms to the Transaction interface
var _ Transaction = &HTTPTransaction{}

// HTTPAttemptHandler is an event handler that will get called each time this transaction is attempted
type HTTPAttemptHandler func(transaction *HTTPTransaction)

// HTTPCompletionHandler is an  event handler that will get called after this transaction has completed
type HTTPCompletionHandler func(transaction *HTTPTransaction, statusCode int, body []byte, err error)

var defaultAttemptHandler = func(transaction *HTTPTransaction) {}
var defaultCompletionHandler = func(transaction *HTTPTransaction, statusCode int, body []byte, err error) {}

func initTransactionExpvars() {
	transactionsInputBytesByEndpoint.Init()
	transactionsConnectionEvents.Init()
	transactionsInputCountByEndpoint.Init()
	transactionsDroppedByEndpoint.Init()
	transactionsRequeuedByEndpoint.Init()
	transactionsRetriedByEndpoint.Init()
	transactionsSuccessByEndpoint.Init()
	transactionsSuccessBytesByEndpoint.Init()
	transactionsErrorsByType.Init()
	transactionsHTTPErrorsByCode.Init()
	transactionsConnectionEvents.Set("DNSSuccess", &connectionDNSSuccess)
	transactionsConnectionEvents.Set("ConnectSuccess", &connectionConnectSuccess)
	transactionsExpvars.Set("InputBytesByEndpoint", &transactionsInputBytesByEndpoint)
	transactionsExpvars.Set("ConnectionEvents", &transactionsConnectionEvents)
	transactionsExpvars.Set("InputCountByEndpoint", &transactionsInputCountByEndpoint)
	transactionsExpvars.Set("Dropped", &transactionsDropped)
	transactionsExpvars.Set("DroppedByEndpoint", &transactionsDroppedByEndpoint)
	transactionsExpvars.Set("DroppedOnInput", &transactionsDroppedOnInput)
	transactionsExpvars.Set("Requeued", &transactionsRequeued)
	transactionsExpvars.Set("RequeuedByEndpoint", &transactionsRequeuedByEndpoint)
	transactionsExpvars.Set("Retried", &transactionsRetried)
	transactionsExpvars.Set("RetriedByEndpoint", &transactionsRetriedByEndpoint)
	transactionsExpvars.Set("RetryQueueSize", &transactionsRetryQueueSize)
	transactionsExpvars.Set("SuccessByEndpoint", &transactionsSuccessByEndpoint)
	transactionsExpvars.Set("SuccessBytesByEndpoint", &transactionsSuccessBytesByEndpoint)
	transactionsExpvars.Set("Success", &transactionsSuccess)
	transactionsExpvars.Set("Errors", &transactionsErrors)
	transactionsExpvars.Set("ErrorsByType", &transactionsErrorsByType)
	transactionsErrorsByType.Set("DNSErrors", &transactionsDNSErrors)
	transactionsErrorsByType.Set("TLSErrors", &transactionsTLSErrors)
	transactionsErrorsByType.Set("ConnectionErrors", &transactionsConnectionErrors)
	transactionsErrorsByType.Set("WroteRequestErrors", &transactionsWroteRequestErrors)
	transactionsErrorsByType.Set("SentRequestErrors", &transactionsSentRequestErrors)
	transactionsExpvars.Set("HTTPErrors", &transactionsHTTPErrors)
	transactionsExpvars.Set("HTTPErrorsByCode", &transactionsHTTPErrorsByCode)
}

// TransactionPriority defines the priority of a transaction
// Transactions with priority `TransactionPriorityNormal` are dropped from the retry queue
// before dropping transactions with priority `TransactionPriorityHigh`.
type TransactionPriority int

const (
	// TransactionPriorityNormal defines a transaction with a normal priority
	TransactionPriorityNormal TransactionPriority = 0

	// TransactionPriorityHigh defines a transaction with an high priority
	TransactionPriorityHigh TransactionPriority = 1
)

// HTTPTransaction represents one Payload for one Endpoint on one Domain.
type HTTPTransaction struct {
	// Domain represents the domain target by the HTTPTransaction.
	Domain string
	// Endpoint is the API Endpoint used by the HTTPTransaction.
	Endpoint endpoint
	// Headers are the HTTP headers used by the HTTPTransaction.
	Headers http.Header
	// Payload is the content delivered to the backend.
	Payload *[]byte
	// ErrorCount is the number of times this HTTPTransaction failed to be processed.
	ErrorCount int

	createdAt time.Time
	// retryable indicates whether this transaction can be retried
	retryable bool

	// attemptHandler will be called with a transaction before the attempting to send the request
	attemptHandler HTTPAttemptHandler
	// completionHandler will be called with a transaction after it has been successfully sent
	completionHandler HTTPCompletionHandler

	priority TransactionPriority
}

// Transaction represents the task to process for a Worker.
type Transaction interface {
	Process(ctx context.Context, client *http.Client) error
	GetCreatedAt() time.Time
	GetTarget() string
	GetPriority() TransactionPriority
	GetEndpointName() string
	GetPayloadSize() int

	// This method serializes the transaction to `TransactionsSerializer`.
	// It forces a new implementation of `Transaction` to define how to
	// serialize the transaction to `TransactionsSerializer` as a `Transaction`
	// must be serializable in domainForwarder.
	SerializeTo(*TransactionsSerializer) error
}

// NewHTTPTransaction returns a new HTTPTransaction.
func NewHTTPTransaction() *HTTPTransaction {
	tr := &HTTPTransaction{
		createdAt:  time.Now(),
		ErrorCount: 0,
		retryable:  true,
		Headers:    make(http.Header),
	}
	tr.setDefaultHandlers()
	return tr
}

func (t *HTTPTransaction) setDefaultHandlers() {
	t.attemptHandler = defaultAttemptHandler
	t.completionHandler = defaultCompletionHandler
}

// GetCreatedAt returns the creation time of the HTTPTransaction.
func (t *HTTPTransaction) GetCreatedAt() time.Time {
	return t.createdAt
}

// GetTarget return the url used by the transaction
func (t *HTTPTransaction) GetTarget() string {
	url := t.Domain + t.Endpoint.route
	return httputils.SanitizeURL(url) // sanitized url that can be logged
}

// GetPriority returns the priority
func (t *HTTPTransaction) GetPriority() TransactionPriority {
	return t.priority
}

// GetEndpointName returns the name of the endpoint used by the transaction
func (t *HTTPTransaction) GetEndpointName() string {
	return t.Endpoint.name
}

// GetPayloadSize returns the size of the payload.
func (t *HTTPTransaction) GetPayloadSize() int {
	if t.Payload != nil {
		return len(*t.Payload)
	}

	return 0
}

// Process sends the Payload of the transaction to the right Endpoint and Domain.
func (t *HTTPTransaction) Process(ctx context.Context, client *http.Client) error {
	t.attemptHandler(t)

	statusCode, body, err := t.internalProcess(ctx, client)

	if err == nil || !t.retryable {
		t.completionHandler(t, statusCode, body, err)
	}

	// If the txn is retryable, return the error (if present) to the worker to allow it to be retried
	// Otherwise, return nil so the txn won't be retried.
	if t.retryable {
		return err
	}

	return nil
}

// internalProcess does the  work of actually sending the http request to the specified domain
// This will return  (http status code, response body, error).
func (t *HTTPTransaction) internalProcess(ctx context.Context, client *http.Client) (int, []byte, error) {
	reader := bytes.NewReader(*t.Payload)
	url := t.Domain + t.Endpoint.route
	transactionEndpointName := t.GetEndpointName()
	logURL := httputils.SanitizeURL(url) // sanitized url that can be logged

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Errorf("Could not create request for transaction to invalid URL %q (dropping transaction): %s", logURL, err)
		transactionsErrors.Add(1)
		tlmTxErrors.Inc(t.Domain, transactionEndpointName, "invalid_request")
		transactionsSentRequestErrors.Add(1)
		return 0, nil, nil
	}
	req = req.WithContext(ctx)
	req.Header = t.Headers
	resp, err := client.Do(req)

	if err != nil {
		// Do not requeue transaction if that one was canceled
		if ctx.Err() == context.Canceled {
			return 0, nil, nil
		}
		t.ErrorCount++
		transactionsErrors.Add(1)
		tlmTxErrors.Inc(t.Domain, transactionEndpointName, "cant_send")
		return 0, nil, fmt.Errorf("error while sending transaction, rescheduling it: %s", httputils.SanitizeURL(err.Error()))
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Fail to read the response Body: %s", err)
		return 0, nil, err
	}

	if resp.StatusCode >= 400 {
		statusCode := strconv.Itoa(resp.StatusCode)
		var codeCount *expvar.Int
		if count := transactionsHTTPErrorsByCode.Get(statusCode); count == nil {
			codeCount = &expvar.Int{}
			transactionsHTTPErrorsByCode.Set(statusCode, codeCount)
		} else {
			codeCount = count.(*expvar.Int)
		}
		codeCount.Add(1)
		transactionsHTTPErrors.Add(1)
		tlmTxHTTPErrors.Inc(t.Domain, transactionEndpointName, statusCode)
	}

	if resp.StatusCode == 400 || resp.StatusCode == 404 || resp.StatusCode == 413 {
		log.Errorf("Error code %q received while sending transaction to %q: %s, dropping it", resp.Status, logURL, string(body))
		transactionsDroppedByEndpoint.Add(transactionEndpointName, 1)
		transactionsDropped.Add(1)
		tlmTxDropped.Inc(t.Domain, transactionEndpointName)
		return resp.StatusCode, body, nil
	} else if resp.StatusCode == 403 {
		log.Errorf("API Key invalid, dropping transaction for %s", logURL)
		transactionsDroppedByEndpoint.Add(transactionEndpointName, 1)
		transactionsDropped.Add(1)
		tlmTxDropped.Inc(t.Domain, transactionEndpointName)
		return resp.StatusCode, body, nil
	} else if resp.StatusCode > 400 {
		t.ErrorCount++
		transactionsErrors.Add(1)
		tlmTxErrors.Inc(t.Domain, transactionEndpointName, "gt_400")
		return resp.StatusCode, body, fmt.Errorf("error %q while sending transaction to %q, rescheduling it", resp.Status, logURL)
	}

	tlmTxSuccessCount.Inc(t.Domain, transactionEndpointName)
	tlmTxSuccessBytes.Add(float64(t.GetPayloadSize()), t.Domain, transactionEndpointName)
	transactionsSuccessByEndpoint.Add(transactionEndpointName, 1)
	transactionsSuccessBytesByEndpoint.Add(transactionEndpointName, int64(t.GetPayloadSize()))
	transactionsSuccess.Add(1)

	loggingFrequency := config.Datadog.GetInt64("logging_frequency")

	if transactionsSuccess.Value() == 1 {
		log.Infof("Successfully posted payload to %q, the agent will only log transaction success every %d transactions", logURL, loggingFrequency)
		log.Tracef("Url: %q payload: %s", logURL, string(body))
		return resp.StatusCode, body, nil
	}
	if transactionsSuccess.Value()%loggingFrequency == 0 {
		log.Infof("Successfully posted payload to %q", logURL)
		log.Tracef("Url: %q payload: %s", logURL, string(body))
		return resp.StatusCode, body, nil
	}
	log.Tracef("Successfully posted payload to %q: %s", logURL, string(body))
	return resp.StatusCode, body, nil
}

// SerializeTo serializes the transaction using TransactionsSerializer
func (t *HTTPTransaction) SerializeTo(serializer *TransactionsSerializer) error {
	return serializer.Add(t)
}
