// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

var (
	// Generators generators list
	Generators = make(map[string]Generator, 4)

	// DefParser default parser
	DefParser Parser
)

// Generator list
var (
	GeneratorGin        = "gin"
	GeneratorChi        = "chi"
	GeneratorMux        = "mux"
	GeneratorHttpRouter = "httprouter"
)

// GenOpts generator options
type GenOpts struct {
	Name    string
	OutPath string
}

// TagMir mir tag's info
type TagMir struct {
	// TODO
}

// Parser parse entries
type Parser interface {
	Name() string
	Parse(entries []interface{}) ([]*TagMir, error)
}

// Generator generate interface code for engine
type Generator interface {
	Name() string
	Generate([]*TagMir, *GenOpts) error
}

// Register generator
func Register(gs ...Generator) {
	for _, g := range gs {
		if g != nil && g.Name() != "" {
			Generators[g.Name()] = g
		}
	}
}

// setDefParser set default parser
func SetDefParser(p Parser) {
	DefParser = p
}
