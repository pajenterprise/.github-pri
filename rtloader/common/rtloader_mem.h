// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019 Datadog, Inc.
#ifndef DATADOG_AGENT_RTLOADER_MEM_H
#define DATADOG_AGENT_RTLOADER_MEM_H

/*! \file rtloader_mem.h
    \brief RtLoader memory wrapper header file.

    The prototypes here defined provide functions to allocate and free memory.
    The goal is to allow us to track allocations if desired.
*/

#include "rtloader_types.h"

#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/*! \fn void _set_memory_tracker_cb(cb_memory_tracker_t cb)
    \brief Sets a callback to be used by rtloader to add memory tracking stats.
    \param object A function pointer to the callback function.

    The callback is expected to be provided by the rtloader caller - in go-context: CGO.
    This function is thread unsafe, be sure to call it early on before multiple threads
    may start using the allocator.
*/
void _set_memory_tracker_cb(cb_memory_tracker_t);

/*! \fn void *_malloc(size_t sz)
    \brief Basic malloc wrapper that will also keep memory stats if enabled.
    \param sz the number of bytes to allocate.

    This function is thread unsafe in its access to the memory tracker. Onle, use this
    logic once the memory tracker has be set (or tracking remains disabled).
*/
void *_malloc(size_t sz);

/*! \fn void _free(void *ptr)
    \brief Basic free wrapper that will also keep memory stats if enabled.
    \param ptr the pointer to the heap region you wish to free.

    This function is thread unsafe in its access to the memory tracker. Onle, use this
    logic once the memory tracker has be set (or tracking remains disabled).
*/
void _free(void *ptr);

#ifdef __cplusplus
#    ifdef _WIN32
#        define __THROW
#    elif __APPLE__
#        define __THROW
#    endif

char *strdupe(const char *s1) __THROW;
#else
char *strdupe(const char *s1);
#endif

#ifdef __cplusplus
}
#endif

#endif
