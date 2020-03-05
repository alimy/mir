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

// generatorChi generator for go-chi
type generatorChi struct{}

// Name name of generator
func (generatorChi) Name() string {
	return core.GeneratorChi
}

// Generate generate interface code
func (generatorChi) Generate(ds core.Descriptors, opts *core.Options) error {
	return generate(ds, opts)
}
