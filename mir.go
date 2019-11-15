// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

import (
	"errors"
	"fmt"
)

var (
	generators        = make(map[string]Generator, 4)
	parser     Parser = parserStructTag{}
)

func init() {
	Register(generatorGin{},
		generatorChi{},
		generatorMux{},
		generatorHttpRouter{})
}

// SetTag set custom mir's struct tag name(eg: mir)
func SetTag(name string) {
	if name != "" {
		tagName = name
	}
}

// Register generator
func Register(gs ...Generator) {
	for _, g := range gs {
		if g != nil && g.Name() != "" {
			generators[g.Name()] = g
		}
	}
}

// Generate generate interface code
func Generate(entries []interface{}, opts *GenOpts) error {
	if opts == nil {
		return errors.New("options is nil")
	}
	generator, exist := generators[opts.Name]
	if !exist {
		return fmt.Errorf("unknow generators that name %s", opts.Name)
	}
	mirTags, err := parser.Parse(entries)
	if err != nil {
		return generator.Generate(mirTags, opts)
	}
	return err
}
