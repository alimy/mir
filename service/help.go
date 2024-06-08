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
