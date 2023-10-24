// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.
// +build windows

package system

import (
	"fmt"
	"runtime"

	core "github.com/DataDog/datadog-agent/pkg/collector/corechecks"
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/DataDog/datadog-agent/pkg/util/winutil"

	"github.com/DataDog/datadog-agent/pkg/aggregator"
)

// For testing purpose
var virtualMemory = winutil.VirtualMemory
var swapMemory = winutil.SwapMemory
var runtimeOS = runtime.GOOS

// MemoryCheck doesn't need additional fields
type MemoryCheck struct {
	core.CheckBase
}

const mbSize float64 = 1024 * 1024

// Run executes the check
func (c *MemoryCheck) Run() error {
	sender, err := aggregator.GetSender(c.ID())
	if err != nil {
		return err
	}

	v, errVirt := virtualMemory()
	if errVirt == nil {
		sender.Gauge("system.mem.total", float64(v.Total)/mbSize, "", nil)
		sender.Gauge("system.mem.free", float64(v.Available)/mbSize, "", nil)
		sender.Gauge("system.mem.usable", float64(v.Available)/mbSize, "", nil)
		sender.Gauge("system.mem.used", float64(v.Total-v.Available)/mbSize, "", nil)
		sender.Gauge("system.mem.pct_usable", float64(100-v.UsedPercent)/100, "", nil)
	} else {
		log.Errorf("system.MemoryCheck: could not retrieve virtual memory stats: %s", errVirt)
	}

	s, errSwap := swapMemory()
	if errSwap == nil {
		sender.Gauge("system.swap.total", float64(s.Total)/mbSize, "", nil)
		sender.Gauge("system.swap.free", float64(s.Free)/mbSize, "", nil)
		sender.Gauge("system.swap.used", float64(s.Used)/mbSize, "", nil)
		sender.Gauge("system.swap.pct_free", float64(100-s.UsedPercent)/100, "", nil)
		sender.Gauge("system.mem.cached", float64(s.Cached), "", nil)
		// Committed memory is physical memory for which space has been
		// reserved on the disk paging file in case it must be written
		// back to disk
		sender.Gauge("system.mem.committed", float64(s.Committed), "", nil)
		// physical memory used by the operating system, for objects
		// that can be written to disk when they are not being used
		sender.Gauge("system.mem.paged", float64(s.Paged), "", nil)
		// physical memory used by the operating system for objects that
		// cannot be written to disk, but must remain in physical memory
		// as long as they are allocated.
		sender.Gauge("system.mem.nonpaged", float64(s.Nonpaged), "", nil)
	} else {
		log.Errorf("system.MemoryCheck: could not retrieve swap memory stats: %s", errSwap)
	}

	if errVirt != nil && errSwap != nil {
		return fmt.Errorf("failed to gather any memory information")
	}

	sender.Commit()
	return nil
}
