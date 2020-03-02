// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"errors"

	"github.com/alimy/mir/v2/core"
)

// generatorGin generator for Gin
type generatorGin struct{}

// Name name of generator
func (generatorGin) Name() string {
	return core.EngineGin
}

// Generate generate interface code
func (generatorGin) Generate(entries []*core.TagMir, opts *core.GenOpts) error {
	// TODO
	return errors.New("not ready")
}
