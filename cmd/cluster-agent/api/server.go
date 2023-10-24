// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017 Datadog, Inc.

/*
Package api implements the agent IPC api. Using HTTP
calls, it's possible to communicate with the agent,
sending commands and receiving infos.
*/
package api

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	stdLog "log"
	"net"
	"net/http"

	"github.com/DataDog/datadog-agent/cmd/cluster-agent/api/agent"
	"github.com/DataDog/datadog-agent/pkg/api/security"
	"github.com/DataDog/datadog-agent/pkg/api/util"
	"github.com/DataDog/datadog-agent/pkg/config"
	"github.com/gorilla/mux"
)

var (
	listener net.Listener
)

// StartServer creates the router and starts the HTTP server
func StartServer() error {
	// create the root HTTP router
	r := mux.NewRouter()

	// IPC REST API server
	agent.SetupHandlers(r)

	// get the transport we're going to use under HTTP
	var err error
	listener, err = getListener()
	if err != nil {
		// we use the listener to handle commands for the agent, there's
		// no way we can recover from this error
		return fmt.Errorf("Unable to create the api server: %v", err)
	}
	util.SetAuthToken()

	// create cert
	hosts := []string{"127.0.0.1", "localhost"}
	_, rootCertPEM, rootKey, err := security.GenerateRootCert(hosts, 2048)
	if err != nil {
		return fmt.Errorf("unable to start TLS server")
	}

	// PEM encode the private key
	rootKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
	})

	// Create a TLS cert using the private key and certificate
	rootTLSCert, err := tls.X509KeyPair(rootCertPEM, rootKeyPEM)
	if err != nil {
		return fmt.Errorf("invalid key pair: %v", err)
	}

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{rootTLSCert},
	}

	srv := &http.Server{
		Handler:   r,
		ErrorLog:  stdLog.New(&config.ErrorLogWriter{}, "", 0), // log errors to seelog
		TLSConfig: &tlsConfig,
	}
	tlsListener := tls.NewListener(listener, &tlsConfig)
	fmt.Println(tlsListener)
	go srv.Serve(listener)
	return nil

}

// StopServer closes the connection and the server
// stops listening to new commands.
func StopServer() {
	if listener != nil {
		listener.Close()
	}
}
