// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

// Launch start task interface
type Launch interface {
	Go(func())
}

// Runtime service runtime interface
type Runtime interface {
	Start(Launch)
	Stop()
}
