// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"context"
	"log"
)

const (
	// run mode list
	InSerialMode RunMode = iota
	InConcurrentMode
	InSerialDebugMode
	InConcurrentDebugMode

	// options key list
	OptSinkPath   = "sinkPath"
	OptDefaultTag = "defaultTag"

	// generator Names
	GeneratorGin        = "gin"
	GeneratorChi        = "chi"
	GeneratorMux        = "mux"
	GeneratorEcho       = "echo"
	GeneratorIris       = "iris"
	GeneratorMacaron    = "macaron"
	GeneratorHttpRouter = "httprouter"

	// parser Names
	ParserStructTag = "structTag"
)

var (
	// generators generator list
	generators = make(map[string]Generator, 4)

	// parsers parser list
	parsers = make(map[string]Parser, 1)

	// InDebug whether in debug mode
	InDebug bool
)

// RunMode indicate process mode (InSerialMode | InSerialDebugMode | InConcurrentMode | InConcurrentDebugMode)
type RunMode uint8

// Opts use for generator or parser init
type InitOpts = map[string]string

// Options generator options
type Options struct {
	// RunMode set run mode (InSerialMode Or InConcurrentMode).
	// Default is InSerialMode if not set explicit.
	RunMode       RunMode
	GeneratorName string
	ParserName    string
	GeneratorOpts InitOpts
	ParserOpts    InitOpts
}

// Crate component common info
type Crate interface {
	Name() string
	Init(opts InitOpts) error
}

// Parser parse entries
type Parser interface {
	Crate
	Parse(entries []interface{}) (Descriptors, error)
	ParseContext(ctx MirCtx, entries []interface{})
	Clone() Parser
}

// Generator generate interface code for engine
type Generator interface {
	Crate
	Generate(Descriptors) error
	GenerateContext(ctx MirCtx)
	Clone() Generator
}

// MirCtx mir's concurrent parser/generator context
type MirCtx interface {
	context.Context
	Cancel(err error)
	ParserDone()
	GeneratorDone()
	Wait() error
	Capcity() int
	Pipe() (<-chan *IfaceDescriptor, chan<- *IfaceDescriptor)
}

func (m RunMode) String() string {
	res := "not support mode"
	switch m {
	case InSerialMode:
		res = "serial mode"
	case InSerialDebugMode:
		res = "serial debug mode"
	case InConcurrentMode:
		res = "concurrent mode"
	case InConcurrentDebugMode:
		res = "concurrent debug mode"
	}
	return res
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

// Logus print log info
func Logus(format string, v ...interface{}) {
	if InDebug {
		log.Printf("[mir] "+format, v...)
	}
}
