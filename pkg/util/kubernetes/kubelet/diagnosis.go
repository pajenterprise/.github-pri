// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017 Datadog, Inc.

// +build docker
// +build kubelet

package kubelet

import (
	"github.com/DataDog/datadog-agent/pkg/diagnose/diagnosis"
	log "github.com/cihub/seelog"
)

func init() {
	diagnosis.Register("Kubelet availability", diagnose)
}

// diagnose the API server availability
func diagnose() error {
	_, err := GetKubeUtil()
	if err != nil {
		log.Error(err)
	}
	return err
}
