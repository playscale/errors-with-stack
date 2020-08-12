// Copyright (c) 2020 Playscale PTE LTD. All rights reserved.
// Licensed under the MIT license. See LICENSE file included with
// this project for details.

package errors

import (
	goerrors "errors"
	"runtime/debug"
	"sync"
)

var useStackTracesMutex = sync.RWMutex{}
var useStackTraces = true

// New will create a new, regular, go Error with your message and a
// stack trace appended. New(msg + "\nStack Trace:\n" + stackTrace)
func New(msg string) error {
	useStackTracesMutex.RLock()
	shouldPrint := useStackTraces
	useStackTracesMutex.RUnlock()
	if !shouldPrint {
		return goerrors.New(msg)
	}

	stackTrace := string(debug.Stack())

	result := goerrors.New(msg + "\nStack Trace:\n" + stackTrace)
	return result
}

// EnableStackTraces does what it says on the tin, and does so in a
// thread-safe manner. Any calls to errors.New(message) after
// EnableStackTraces() will append a formatted stack trace to the error
// message.
func EnableStackTraces() {
	useStackTracesMutex.Lock()
	defer useStackTracesMutex.Unlock()
	useStackTraces = true
}

// DisableStackTraces does just that. It is safe to call this from
// multiple threads. Any calls to errors.New(message) after
// DisableStackTraces() will simply return a plain old Go standard
// library call to errors.New() as if this package were not even in
// place.
func DisableStackTraces() {
	useStackTracesMutex.Lock()
	defer useStackTracesMutex.Unlock()
	useStackTraces = false
}
