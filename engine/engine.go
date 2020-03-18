// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/alimy/mir/v2/core"

	_ "github.com/alimy/mir/v2/internal"
)

// Generate generate interface code
func Generate(entries []interface{}, opts *core.Options) (err error) {
	if opts == nil {
		return errors.New("options is nil")
	}

	if len(entries) == 0 {
		return errors.New("entries is empty")
	}

	p := core.ParserByName(opts.ParserName)
	// use default parser when not set parser name from options
	if p == nil {
		p = core.DefaultParser()
	}
	if err = p.Init(opts.ParserOpts); err != nil {
		return
	}

	g := core.GeneratorByName(opts.GeneratorName)
	if g == nil {
		return fmt.Errorf("unknow generators that name %s", opts.GeneratorName)
	}
	if err = g.Init(opts.GeneratorOpts); err != nil {
		return
	}

	switch opts.RunMode {
	case core.InSerialDebugMode:
		core.InDebug = true
		core.Logus("run in %s", opts.RunMode)
		fallthrough
	case core.InSerialMode:
		err = doInSerial(p, g, entries)
	case core.InConcurrentDebugMode:
		core.InDebug = true
		core.Logus("run in %s", opts.RunMode)
		fallthrough
	case core.InConcurrentMode:
		err = doInConcurrent(p, g, entries)
	}
	return err
}

func doInSerial(p core.Parser, g core.Generator, entries []interface{}) error {
	descriptors, err := p.Parse(entries)
	if err == nil {
		return g.Generate(descriptors)
	}
	return err
}

func doInConcurrent(p core.Parser, g core.Generator, entries []interface{}) error {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	core.Logus("set GOMAXPROCS: %d", numCPU)

	ctx := core.NewMirCtx(10)

	go p.ParseContext(ctx, entries)
	go g.GenerateContext(ctx)

	select {
	case <-ctx.Done():
		if ctx.IsGeneratorDone() {
			return nil
		}
		return ctx.Err()
	}
}
