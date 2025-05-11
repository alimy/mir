// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"io"
	"os"
)

var (
	_output io.Writer = os.Stdout
)

// SetOutput set log output writer. default is os.Stdout
func SetOutput(w io.Writer) {
	if w != nil {
		_output = w
	}
}
