// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"errors"
	"fmt"

	"github.com/alimy/mir/v2/core"

	_ "github.com/alimy/mir/v2/generator"
	_ "github.com/alimy/mir/v2/parser"
)

// Generate generate interface code
func Generate(entries []interface{}, opts *core.GenOpts) error {
	if opts == nil {
		return errors.New("options is nil")
	}
	g, exist := core.Generators[opts.Name]
	if !exist {
		return fmt.Errorf("unknow generators that name %s", opts.Name)
	}
	if core.DefParser == nil {
		return errors.New("parser is nil")
	}
	mirTags, err := core.DefParser.Parse(entries)
	if err != nil {
		return g.Generate(mirTags, opts)
	}
	return err
}
