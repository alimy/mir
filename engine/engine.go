// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"errors"
	"fmt"

	"github.com/alimy/mir/v2/core"

	_ "github.com/alimy/mir/v2/internal"
)

// Generate generate interface code
func Generate(entries []interface{}, opts *core.Options) (err error) {
	if opts == nil {
		return errors.New("options is nil")
	}

	p := core.ParserByName(opts.ParserName)
	// use default parser when not set parser name from options
	if p == nil {
		p = core.DefaultParser()
	}
	if err = p.Init(opts.ParserOpts); err != nil {
		return
	}

	g := core.GeneratorByName(opts.GeneratorName)
	if g == nil {
		return fmt.Errorf("unknow generators that name %s", opts.GeneratorName)
	}
	if err = g.Init(opts.GeneratorOpts); err != nil {
		return
	}

	if mirTags, err := p.Parse(entries); err == nil {
		return g.Generate(mirTags)
	}
	return err
}
