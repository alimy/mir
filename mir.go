// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

import "errors"

var (
	generators = make(map[string]Generator)
)

// SetTag set custom mir's struct tag name(eg: mir)
func SetTag(name string) {
	if name != "" {
		tagName = name
	}
}

// Register generator
func Register(g Generator) {
	if g != nil {
		generators[g.Name()] = g
	}
}

// Generate generate interface code
func Generate(entries []interface{}, opts *GenOpts) error {
	// TODO
	return errors.New("not ready")
}
