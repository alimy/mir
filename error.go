// Copyright 2022 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

import (
	"errors"
	"fmt"
)

// Error indicator error's wraper
type Error interface {
	error
	StatusCode() int
}

type httpError struct {
	error
	statusCode int
}

func NewError(statusCode int, err error) Error {
	return &httpError{
		error:      err,
		statusCode: statusCode,
	}
}

func Errorf(statusCode int, format string, a ...any) Error {
	return &httpError{
		error:      fmt.Errorf(format, a...),
		statusCode: statusCode,
	}
}

func Errorln(statusCode int, text string) Error {
	return &httpError{
		error:      errors.New(text),
		statusCode: statusCode,
	}
}

func (e *httpError) StatusCode() int {
	return e.statusCode
}
