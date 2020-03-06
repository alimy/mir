// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

// options key list
const (
	OptSinkPath   = "sinkPath"
	OptDefaultTag = "defaultTag"
)

var (
	// generators generator list
	generators = make(map[string]Generator, 4)

	// parsers parser list
	parsers = make(map[string]Parser, 1)

	// Generator Names
	GeneratorGin        = "gin"
	GeneratorChi        = "chi"
	GeneratorMux        = "mux"
	GeneratorHttpRouter = "httprouter"

	// Parser Names
	ParserStructTag = "structTag"
)

// Opts use for generator or parser init
type InitOpts = map[string]string

// Options generator options
type Options struct {
	GeneratorName string
	ParserName    string
	GeneratorOpts InitOpts
	ParserOpts    InitOpts
}

// Parser parse entries
type Parser interface {
	Name() string
	Init(opts InitOpts) error
	Parse(entries []interface{}) (Descriptors, error)
}

// Generator generate interface code for engine
type Generator interface {
	Name() string
	Init(opts InitOpts) error
	Generate(Descriptors) error
}

// RegisterGenerators register generators
func RegisterGenerators(gs ...Generator) {
	for _, g := range gs {
		if g != nil && g.Name() != "" {
			generators[g.Name()] = g
		}
	}
}

// RegisterParsers register parsers
func RegisterParsers(ps ...Parser) {
	for _, p := range ps {
		if p != nil && p.Name() != "" {
			parsers[p.Name()] = p
		}
	}
}

// DefaultOptions get a default options
func DefaultOptions() *Options {
	return &Options{
		GeneratorName: GeneratorGin,
		ParserName:    ParserStructTag,
		GeneratorOpts: InitOpts{
			OptSinkPath: "./gen",
		},
	}
}

// GeneratorByName get a generator by name
func GeneratorByName(name string) Generator {
	return generators[name]
}

// DefaultGenerator get a default generator
func DefaultGenerator() Generator {
	return generators[GeneratorGin]
}

// ParserByName get a parser by name
func ParserByName(name string) Parser {
	return parsers[name]
}

// DefaultParser get a default parser
func DefaultParser() Parser {
	return parsers[ParserStructTag]
}
