// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"context"
	"log"

	"github.com/alimy/mir/v4/assert"
)

const (
	// run mode list
	InSerialMode runMode = iota
	InConcurrentMode
	InSerialDebugMode
	InConcurrentDebugMode

	// generator Names
	GeneratorGin        = "gin"
	GeneratorChi        = "chi"
	GeneratorMux        = "mux"
	GeneratorHertz      = "hertz"
	GeneratorEcho       = "echo"
	GeneratorIris       = "iris"
	GeneratorFiber      = "fiber"
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

	// inDebug whether in debug mode
	inDebug bool
)

// runMode indicate process mode (InSerialMode | InSerialDebugMode | InConcurrentMode | InConcurrentDebugMode)
type runMode uint8

// InitOpts use for generator or parser init
type InitOpts struct {
	RunMode           runMode
	GeneratorName     string
	ParserName        string
	SinkPath          string
	DefaultTag        string
	EnginePkgName     string
	EngineImportAlias string
	WatchCtxDone      bool
	NoneQuery         bool
	Cleanup           bool
}

// ParserOpts used for initial parser
type ParserOpts struct {
	EngineInfo   *EngineInfo
	DefaultTag   string
	WatchCtxDone bool
	NoneQuery    bool
}

// GeneratorOpts used for initial generator
type GeneratorOpts struct {
	SinkPath string
	Cleanup  bool
}

// Option pass option to custom run behavior
type Option interface {
	apply(opts *InitOpts)
}

// Options generator options
type Options []Option

// InitOpts return an initOpts instance
func (opts Options) InitOpts() *InitOpts {
	res := defaultInitOpts()
	for _, opt := range opts {
		opt.apply(res)
	}
	return res
}

// ParserOpts return a ParserOpts instance
func (opts *InitOpts) ParserOpts() *ParserOpts {
	return &ParserOpts{
		DefaultTag:   opts.DefaultTag,
		WatchCtxDone: opts.WatchCtxDone,
		NoneQuery:    opts.NoneQuery,
		EngineInfo: &EngineInfo{
			PkgName:     opts.EnginePkgName,
			ImportAlias: opts.EngineImportAlias,
		},
	}
}

// GeneratorOpts return a GeneratorOpts
func (opts *InitOpts) GeneratorOpts() *GeneratorOpts {
	return &GeneratorOpts{
		SinkPath: opts.SinkPath,
		Cleanup:  opts.Cleanup,
	}
}

// optFunc used for convert function to Option interface
type optFunc func(opts *InitOpts)

func (f optFunc) apply(opts *InitOpts) {
	f(opts)
}

// Parser parse entries
type Parser interface {
	Name() string
	Init(opts *ParserOpts) error
	Parse(entries []any) (Descriptors, error)
	ParseContext(ctx MirCtx, entries []any)
	Clone() Parser
}

// Generator generate interface code for engine
type Generator interface {
	Name() string
	Init(opts *GeneratorOpts) error
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

// String runMode describe
func (m runMode) String() string {
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

// RunMode set run mode option
func RunMode(mode runMode) Option {
	return optFunc(func(opts *InitOpts) {
		opts.RunMode = mode
	})
}

// GeneratorName set generator name option
func GeneratorName(name string) Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = name
	})
}

// UseGin use Gin engine generator
func UseGin() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorGin
	})
}

// UseChi use Chi engine generator
func UseChi() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorChi
	})
}

// UseMux use Mux engine generator
func UseMux() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorMux
	})
}

// UseHertz use Hertz engine generator
func UseHertz() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorHertz
	})
}

// UseEcho use Echo engine generator
func UseEcho() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorEcho
	})
}

// UseIris use Iris engine generator
func UseIris() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorIris
	})
}

// UseFiber use Fiber engine generator
func UseFiber() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorFiber
	})
}

// UseMacaron use Macaron engine generator
func UseMacaron() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorMacaron
	})
}

// UseHttpRouter use HttpRouter engine generator
func UseHttpRouter() Option {
	return optFunc(func(opts *InitOpts) {
		opts.GeneratorName = GeneratorHttpRouter
	})
}

// ParserName set parser name option
func ParserName(name string) Option {
	return optFunc(func(opts *InitOpts) {
		opts.ParserName = name
	})
}

// SinkPath set generated code out directory
func SinkPath(path string) Option {
	return optFunc(func(opts *InitOpts) {
		opts.SinkPath = path
	})
}

// AssertType[T] register assert.TypeAssertor for custom T type
func AssertType[T any]() Option {
	return optFunc(func(_opts *InitOpts) {
		assert.RegisterType[T]()
	})
}

// AssertType2[B, R] register assert.TypeAssertor for custom B(Binding) and R(Render) type
func AssertType2[B, R any]() Option {
	return optFunc(func(_opts *InitOpts) {
		assert.RegisterType2[B, R]()
	})
}

// AssertType3[B, P, R] register assert.TypeAssertor for custom B(Binding)/P(Params) and R(Render) type
func AssertType3[B, P, R any]() Option {
	return optFunc(func(_opts *InitOpts) {
		assert.RegisterType3[B, P, R]()
	})
}

// WatchCtxDone set generator whether watch context done when Register Servants in
// generated code. default watch context done.
func WatchCtxDone(enable bool) Option {
	return optFunc(func(opts *InitOpts) {
		opts.WatchCtxDone = enable
	})
}

// Cleanup set generator cleanup out first when re-generate code
func Cleanup(enable bool) Option {
	return optFunc(func(opts *InitOpts) {
		opts.Cleanup = enable
	})
}

// NoneQuery set parser whether parse query
func NoneQuery(enable bool) Option {
	return optFunc(func(opts *InitOpts) {
		opts.NoneQuery = enable
	})
}

// DefaultTag set parser's default struct field tag string key
func DefaultTag(tag string) Option {
	return optFunc(func(opts *InitOpts) {
		opts.DefaultTag = tag
	})
}

// EnginePkgName engine package name
func EnginePkgName(pkgName string) Option {
	return optFunc(func(opts *InitOpts) {
		opts.EnginePkgName = pkgName
	})
}

// EngineImportAlias import package alias name
func EngineImportAlias(name string) Option {
	return optFunc(func(opts *InitOpts) {
		opts.EngineImportAlias = name
	})
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
func Logus(format string, v ...any) {
	if inDebug {
		log.Printf("[mir] "+format, v...)
	}
}

// InitFrom initial from Options and return an InitOpts instance
func InitFrom(opts Options) *InitOpts {
	var initOpts *InitOpts
	if opts == nil {
		initOpts = defaultInitOpts()
	} else {
		initOpts = opts.InitOpts()
	}

	switch initOpts.RunMode {
	case InSerialDebugMode, InConcurrentDebugMode:
		inDebug = true
	default:
		inDebug = false
	}

	return initOpts
}

func defaultInitOpts() *InitOpts {
	return &InitOpts{
		RunMode:       InSerialMode,
		GeneratorName: GeneratorGin,
		ParserName:    ParserStructTag,
		SinkPath:      ".auto",
		DefaultTag:    "mir",
		WatchCtxDone:  true,
		Cleanup:       true,
	}
}
