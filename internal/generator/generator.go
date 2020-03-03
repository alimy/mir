// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"github.com/alimy/mir/v2/core"
)

func init() {
	core.RegisterGenerators(generatorGin{},
		generatorChi{},
		generatorMux{},
		generatorHttpRouter{})
}
