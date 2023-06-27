// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"errors"
	"fmt"
	"runtime"
	"sync"

	"github.com/alimy/mir/v4/core"
	"github.com/alimy/mir/v4/internal"
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

// AddEntries add mir's entry list.
// Deprecated use AddEntry(...) instead.
func AddEntries(entries ...any) {
	mu.Lock()
	defer mu.Unlock()

	addEntries(entries...)
}

// Generate generate interface code from mir's iface entry
func Generate(opts core.Options, entries ...any) (err error) {
	mu.Lock()
	defer mu.Unlock()

	addEntries(entries...)
	if len(mirEntries) == 0 {
		return errors.New("mir entries is empty maybe need add entries first")
	}

	initOpts := core.InitFrom(opts)
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
