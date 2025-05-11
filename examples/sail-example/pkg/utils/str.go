// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"unsafe"
)

// String convert bytes to string
func String(data []byte) (res string) {
	if size := len(data); size > 0 {
		res = unsafe.String(unsafe.SliceData(data), size)
	}
	return
}

// Bytes convert string to []byte
func Bytes(data string) (res []byte) {
	if size := len(data); size > 0 {
		res = unsafe.Slice(unsafe.StringData(data), size)
	} else {
		res = []byte{}
	}
	return
}
