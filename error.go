// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
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

// httpError a http error that implements Error interface.
type httpError struct {
	error
	statusCode int
}

func (e *httpError) StatusCode() int {
	return e.statusCode
}

// NewError create a httpError instance.
func NewError(statusCode int, err error) Error {
	return &httpError{
		error:      err,
		statusCode: statusCode,
	}
}

// Errorf create a httpError instance by format.
func Errorf(statusCode int, format string, a ...any) Error {
	return &httpError{
		error:      fmt.Errorf(format, a...),
		statusCode: statusCode,
	}
}

// Errorf create a httpError instance by text.
func Errorln(statusCode int, text string) Error {
	return &httpError{
		error:      errors.New(text),
		statusCode: statusCode,
	}
}

// HttpStatusCode get http status code from error that implements Error interface.
func HttpStatusCode(err error) (code int, ok bool) {
	for {
		if x, ok := err.(interface{ StatusCode() int }); ok {
			return x.StatusCode(), true
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if err == nil {
					continue
				}
				if code, ok = HttpStatusCode(err); ok {
					return
				}
			}
			return
		default:
			return
		}
	}
}
