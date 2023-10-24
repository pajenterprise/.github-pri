// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

// +build !windows,!android

//go:generate go run ../../pkg/config/render_config.go agent ../../pkg/config/config_template.yaml ./dist/datadog.yaml
//go:generate go run ../../pkg/config/render_config.go network-tracer ../../pkg/config/config_template.yaml ./dist/network-tracer.yaml

package main

import (
	"os"

	"github.com/DataDog/datadog-agent/cmd/agent/app"
	reaper "github.com/ramr/go-reaper"
)

func main() {
	// TODO: docker only
	// Reap orphaned child processes
	go reaper.Reap()

	// Invoke the Agent
	if err := app.AgentCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
