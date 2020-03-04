// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.
// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"errors"

	"github.com/alimy/mir/v2/core"
)

// generatorHttpRouter generator for HttpRouter
type generatorHttpRouter struct{}

// Name name of generator
func (generatorHttpRouter) Name() string {
	return core.GeneratorHttpRouter
}

// Generate generate interface code
func (generatorHttpRouter) Generate(ds core.Descriptors, opts *core.Options) error {
	// TODO
	return errors.New("not ready")
}
