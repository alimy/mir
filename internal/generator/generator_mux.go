// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.
// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"github.com/alimy/mir/v2/core"
)

// generatorMux generator for Mux
type generatorMux struct{}

// Name name of generator
func (generatorMux) Name() string {
	return core.GeneratorMux
}

// Generate generate interface code
func (generatorMux) Generate(ds core.Descriptors, opts *core.Options) error {
	return generate(ds, opts)
}
