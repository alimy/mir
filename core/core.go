// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"log"
	"os"
	"path/filepath"
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

// Options generator options
type Options struct {
	GeneratorName string
	ParserName    string
	OutPath       string
}

// Parser parse entries
type Parser interface {
	Name() string
	Parse(entries []interface{}) (Descriptors, error)
}

// Generator generate interface code for engine
type Generator interface {
	Name() string
	Generate(Descriptors, *Options) error
}

// SinkPath return output path
func (p *Options) SinkPath() string {
	path, err := filepath.EvalSymlinks(p.OutPath)
	if err != nil {
		if os.IsNotExist(err) {
			if !filepath.IsAbs(p.OutPath) {
				cwd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				path = filepath.Join(cwd, p.OutPath)
			} else {
				path = p.OutPath
			}
		} else {
			log.Fatal(err)
		}
	}
	return path
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
		OutPath:       "./gen",
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
