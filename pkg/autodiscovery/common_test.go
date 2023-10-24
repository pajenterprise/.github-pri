// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package autodiscovery

import (
	"github.com/DataDog/datadog-agent/pkg/autodiscovery/integration"
	"github.com/DataDog/datadog-agent/pkg/autodiscovery/listeners"
)

type dummyService struct {
	ID            string
	ADIdentifiers []string
	Hosts         map[string]string
	Ports         []listeners.ContainerPort
	Pid           int
	Hostname      string
	CreationTime  integration.CreationTime
	UnixSockets   []string
}

// GetEntity returns the service entity name
func (s *dummyService) GetEntity() string {
	return s.ID
}

// GetADIdentifiers returns dummy identifiers
func (s *dummyService) GetADIdentifiers() ([]string, error) {
	return s.ADIdentifiers, nil
}

// GetHosts returns dummy hosts
func (s *dummyService) GetHosts() (map[string]string, error) {
	return s.Hosts, nil
}

// GetPorts returns dummy ports
func (s *dummyService) GetPorts() ([]listeners.ContainerPort, error) {
	return s.Ports, nil
}

// GetTags returns mil
func (s *dummyService) GetTags() ([]string, error) {
	return nil, nil
}

// GetPid return a dummy pid
func (s *dummyService) GetPid() (int, error) {
	return s.Pid, nil
}

// GetHostname return a dummy hostname
func (s *dummyService) GetHostname() (string, error) {
	return s.Hostname, nil
}

// GetCreationTime return a dummy creation time
func (s *dummyService) GetCreationTime() integration.CreationTime {
	return s.CreationTime
}

// GetUnixSockets returns a dummy unix sockets list
func (s *dummyService) GetUnixSockets() ([]string, error) {
	return s.UnixSockets, nil
}
