// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"errors"
	"fmt"

	"github.com/alimy/mir/v2/core"

	_ "github.com/alimy/mir/v2/internal/generator"
	_ "github.com/alimy/mir/v2/internal/parser"
)

// Generate generate interface code
func Generate(entries []interface{}, opts *core.Options) error {
	if opts == nil {
		return errors.New("options is nil")
	}

	// just use default parser now
	p := core.DefaultParser()
	if p == nil {
		return errors.New("parser is nil")
	}

	g := core.GeneratorByName(opts.GeneratorName)
	if g == nil {
		return fmt.Errorf("unknow generators that name %s", opts.GeneratorName)
	}

	mirTags, err := p.Parse(entries)
	if err != nil {
		return g.Generate(mirTags, opts)
	}
	return err
}
