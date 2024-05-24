// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import "github.com/sourcegraph/conc"

type Runtime interface {
	Start(wg *conc.WaitGroup)
	Stop()
}
