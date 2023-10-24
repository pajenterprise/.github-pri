// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build apm
// +build !windows
// +build !linux

// linux handled by systemd/upstart

package embed

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DataDog/datadog-agent/pkg/util/executable"
)

const apm_binary_name = "trace-agent" //nolint:golint

func getAPMAgentDefaultBinPath() (string, error) {
	here, _ := executable.Folder()
	binPath := filepath.Join(here, "..", "..", "embedded", "bin", apm_binary_name)
	_, err := os.Stat(binPath)
	if err == nil {
		return binPath, nil
	}
	return binPath, fmt.Errorf("Can't access the default apm binary at %s: %s", binPath, err)
}
