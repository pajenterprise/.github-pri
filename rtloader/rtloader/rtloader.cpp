// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog
// (https://www.datadoghq.com/).
// Copyright 2019-2020 Datadog, Inc.

#include "rtloader.h"
#include "rtloader_mem.h"

void RtLoader::setError(const std::string &msg) const
{
    _errorFlag = true;
    _error = msg;
}

void RtLoader::setError(const char *msg) const
{
    _errorFlag = true;
    _error = msg;
}

const char *RtLoader::getError() const
{
    if (!_errorFlag) {
        // error was already fetched, cleanup
        _error = "";
    } else {
        _errorFlag = false;
    }

    return _error.c_str();
}

bool RtLoader::hasError() const
{
    return _errorFlag;
}

void RtLoader::clearError()
{
    _errorFlag = false;
    _error = "";
}

void RtLoader::free(void *ptr)
{
    if (ptr != NULL) {
        _free(ptr);
    }
}
