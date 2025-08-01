//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package errors

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/sirupsen/logrus"

	entcfg "github.com/weaviate/weaviate/entities/config"
	entsentry "github.com/weaviate/weaviate/entities/sentry"
	"golang.org/x/sync/errgroup"
)

// ErrorGroupWrapper is a custom type that embeds errgroup.Group.
type ErrorGroupWrapper struct {
	*errgroup.Group
	returnError    error
	variables      []interface{}
	logger         logrus.FieldLogger
	deferFunc      func(localVars ...interface{})
	cancelCtx      func()
	routineCounter int
	includeStack   bool
	limitSet       int
}

// NewErrorGroupWrapper creates a new ErrorGroupWrapper.
func NewErrorGroupWrapper(logger logrus.FieldLogger, vars ...interface{}) *ErrorGroupWrapper {
	egw := &ErrorGroupWrapper{
		Group:       new(errgroup.Group),
		returnError: nil,
		variables:   vars,
		logger:      logger,

		// this dummy func makes it safe to call cancelCtx even if a wrapper without a
		// context is used. Avoids a nil check later on.
		cancelCtx: func() {},
	}
	egw.setDeferFunc()

	if entcfg.Enabled(os.Getenv("LOG_STACK_TRACE_ON_ERROR_GROUP")) {
		egw.includeStack = true
	}
	return egw
}

// NewErrorGroupWithContextWrapper creates a new ErrorGroupWrapper
func NewErrorGroupWithContextWrapper(logger logrus.FieldLogger, ctx context.Context, vars ...interface{}) (*ErrorGroupWrapper, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	eg, ctx := errgroup.WithContext(ctx)
	egw := &ErrorGroupWrapper{
		Group:       eg,
		returnError: nil,
		variables:   vars,
		logger:      logger,
		cancelCtx:   cancel,
	}
	egw.setDeferFunc()

	if entcfg.Enabled(os.Getenv("LOG_STACK_TRACE_ON_ERROR_GROUP")) {
		egw.includeStack = true
	}

	return egw, ctx
}

func (egw *ErrorGroupWrapper) setDeferFunc() {
	disable := entcfg.Enabled(os.Getenv("DISABLE_RECOVERY_ON_PANIC"))
	if !disable {
		egw.deferFunc = func(localVars ...interface{}) {
			if r := recover(); r != nil {
				entsentry.Recover(r)
				egw.logger.WithField("panic", r).Errorf("Recovered from panic: %v, local variables %v, additional localVars %v\n", r, localVars, egw.variables)
				debug.PrintStack()
				egw.returnError = fmt.Errorf("panic occurred: %v", r)
				egw.cancelCtx()
			}
		}
	} else {
		egw.deferFunc = func(localVars ...interface{}) {}
	}
}

// Go overrides the Go method to add panic recovery logic.
func (egw *ErrorGroupWrapper) Go(f func() error, localVars ...interface{}) {
	egw.Group.Go(func() error {
		defer egw.deferFunc(localVars)
		return f()
	})
	egw.routineCounter++
}

// SetLimit overrides the SetLimit method to set a limit on the number of
// goroutines and track what's set.
func (egw *ErrorGroupWrapper) SetLimit(limit int) {
	egw.Group.SetLimit(limit)
	egw.limitSet = limit
}

// Wait waits for all goroutines to finish and returns the first non-nil error.
func (egw *ErrorGroupWrapper) Wait() error {
	logBase := egw.logger.WithFields(logrus.Fields{
		"action":     "error_group_wait_initiated",
		"jobs_count": egw.routineCounter,
		"limit":      egw.limitSet,
	})

	if egw.includeStack {
		stackBuf := make([]byte, 4096)
		n := runtime.Stack(stackBuf, false)
		stackBuf = stackBuf[:n]

		logBase = logBase.WithField("stack", string(stackBuf))
	}

	logBase.Debugf("Waiting for %d jobs to finish with limit %d", egw.routineCounter, egw.limitSet)

	if err := egw.Group.Wait(); err != nil {
		return err
	}
	return egw.returnError
}
