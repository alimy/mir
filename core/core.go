// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"context"
	"errors"
	"log"
	"sync"
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
	GoParse(ctx MirCtx, entries []interface{})
	Clone() Parser
}

// Generator generate interface code for engine
type Generator interface {
	Crate
	Generate(Descriptors) error
	GoGenerate(ctx MirCtx)
	Clone() Generator
}

// MirCtx mir's concurrent parser/generator context
type MirCtx interface {
	context.Context
	Cancel(err error)
	ParserDone()
	GeneratorDone()
	IsGeneratorDone() bool
	ChanCapcity() int
	Pipe() (<-chan *IfaceDescriptor, chan<- *IfaceDescriptor)
}

type mirCtx struct {
	context.Context

	mu           *sync.Mutex
	err          error
	chanCapacity int
	ifaceChan    chan *IfaceDescriptor
	cancelFunc   context.CancelFunc
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

// Err return cancel error
func (c *mirCtx) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.err
}

// ChanCapacity return ifaceChan's capacity
func (c *mirCtx) ChanCapcity() int {
	return c.chanCapacity
}

// Cancel cancel mir's process logic with an error
func (c *mirCtx) Cancel(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.err = err
	c.cancelFunc()
}

// IsGeneratorDone whether generator process done
func (c *mirCtx) IsGeneratorDone() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return errors.Is(c.err, errGeneratorDone{})
}

// GeneratorDone mark generator process done
func (c *mirCtx) GeneratorDone() {
	c.Cancel(errGeneratorDone{})
}

// ParserDone mark parser process  done
func (c *mirCtx) ParserDone() {
	close(c.ifaceChan)
}

// Pipe return source/sink chan *IfaceDescriptor
func (c *mirCtx) Pipe() (<-chan *IfaceDescriptor, chan<- *IfaceDescriptor) {
	return c.ifaceChan, c.ifaceChan
}

// NewMirCtx return a new mir's context instance
func NewMirCtx(capcity int) MirCtx {
	ctx := &mirCtx{
		mu:           &sync.Mutex{},
		chanCapacity: capcity,
		ifaceChan:    make(chan *IfaceDescriptor, capcity),
	}
	ctx.Context, ctx.cancelFunc = context.WithCancel(context.Background())
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

// Logus print log info
func Logus(format string, v ...interface{}) {
	if InDebug {
		log.Printf("[mir] "+format, v...)
	}
}
