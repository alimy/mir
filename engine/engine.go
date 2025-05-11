// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"sync"

	"github.com/alimy/mir/v5/core"
	enginternal "github.com/alimy/mir/v5/engine/internal"
	"github.com/alimy/mir/v5/internal"
)

var (
	mu         = &sync.Mutex{}
	mirEntries = make([]any, 0, 8)
)

// AddEntry add mir's entry list.
func AddEntry(entries ...any) {
	mu.Lock()
	defer mu.Unlock()

	addEntries(entries...)
}

// Entry add mir's entry by a type
func Entry[T any]() {
	mu.Lock()
	defer mu.Unlock()

	addEntries(new(T))
}

// Generate generate interface code from mir's iface entry
func Generate(opts ...core.Option) (err error) {
	mu.Lock()
	defer mu.Unlock()

	initOpts := core.InitFrom(opts)

	if initOpts.UseLoad {
		return generate(initOpts)
	}

	addEntries(initOpts.Entries...)
	if len(mirEntries) == 0 {
		return errors.New("mir entries is empty maybe need add entries first")
	}

	p := core.ParserByName(initOpts.ParserName)
	// use default parser when not set parser name from options
	if p == nil {
		p = core.DefaultParser()
	}
	if err = p.Init(initOpts.ParserOpts()); err != nil {
		return
	}

	g := core.GeneratorByName(initOpts.GeneratorName)
	if g == nil {
		return fmt.Errorf("unknow or yet not supported generator that name of %s", initOpts.GeneratorName)
	}
	if err = g.Init(initOpts.GeneratorOpts()); err != nil {
		return
	}

	core.Logus("run in %s", initOpts.RunMode)
	switch initOpts.RunMode {
	case core.InSerialDebugMode, core.InSerialMode:
		err = doInSerial(p, g, mirEntries)
	case core.InConcurrentDebugMode, core.InConcurrentMode:
		err = doInConcurrent(p, g, mirEntries)
	}
	return err
}

func generate(opts *core.InitOpts) error {
	opts.UseLoad = false
	conf, err := json.MarshalIndent(opts, "", "    ")
	if err != nil {
		return err
	}
	return load(opts.InDebug, opts.GeneratorName, opts.SchemaPath, string(conf))
}

func load(indebug bool, generatorName string, schemaPath []string, conf string) error {
	assertTypeSpec, assertTypeImports := core.AssertTypeSpec(generatorName)
	loader := enginternal.NewLoader(&enginternal.Config{
		InDebug:           indebug,
		InitOpts:          conf,
		SchemaPath:        schemaPath,
		BuildFlags:        []string{"-tags", "mir"},
		AssertTypeImports: assertTypeImports,
		AssertTypeSpec:    assertTypeSpec,
	})
	return loader.Load()
}

func addEntries(entries ...any) {
	mirEntries = append(mirEntries, entries...)
}

func doInSerial(p core.Parser, g core.Generator, entries []any) error {
	descriptors, err := p.Parse(entries)
	if err == nil {
		return g.Generate(descriptors)
	}
	return err
}

func doInConcurrent(p core.Parser, g core.Generator, entries []any) error {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	core.Logus("set GOMAXPROCS: %d", numCPU)

	ctx := internal.NewMirCtx(16)

	go p.ParseContext(ctx, entries)
	go g.GenerateContext(ctx)

	return ctx.Wait()
}
