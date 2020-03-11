// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"context"
	"errors"
	"sync"
)

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

	// generator Names
	GeneratorGin        = "gin"
	GeneratorChi        = "chi"
	GeneratorMux        = "mux"
	GeneratorHttpRouter = "httprouter"

	// parser Names
	ParserStructTag = "structTag"

	// run mode list
	InSerialMode     = false
	InConcurrentMode = true
)

// RunMode indicate process mode (InSerialMode Or InConcurrentMode)
type RunMode = bool

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

// MirCtx mir's concurrent parser/generator context
type MirCtx struct {
	context.Context
	cancelFunc context.CancelFunc
	IfaceChan  chan *IfaceDescriptor
	mu         sync.Mutex
	err        error
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
	GoParse(ctx *MirCtx, entries []interface{})
	Clone() Parser
}

// Generator generate interface code for engine
type Generator interface {
	Crate
	Generate(Descriptors) error
	GoGenerate(ctx *MirCtx)
	Clone() Generator
}

// errGeneratorDone indicate generator process done
type errGeneratorDone struct{}

func (errGeneratorDone) Error() string {
	return "generator process done"
}

func (errGeneratorDone) Is(err error) bool {
	_, ok := err.(errGeneratorDone)
	return ok
}

// Err return cancel error
func (c *MirCtx) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.err
}

// Cancel cancel mir's process logic with an error
func (c *MirCtx) Cancel(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.err = err
	c.cancelFunc()
}

// IsGeneratorDone whether generator process done
func (c *MirCtx) IsGeneratorDone() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return errors.Is(c.err, errGeneratorDone{})
}

// GeneratorDone mark generator process done
func (c *MirCtx) GeneratorDone() {
	c.Cancel(errGeneratorDone{})
}

// ParserDone mark parser process  done
func (c *MirCtx) ParserDone() {
	close(c.IfaceChan)
}

// NewMirCtx return a new *MirCtx instance
func NewMirCtx(capcity int) *MirCtx {
	ctx := &MirCtx{
		IfaceChan: make(chan *IfaceDescriptor, capcity),
		mu:        sync.Mutex{},
	}
	ctx.Context, ctx.cancelFunc = context.WithCancel(ctx)
	return ctx
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
