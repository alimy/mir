// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"errors"
	"fmt"

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
	case core.InSerialMode:
		err = doInSerial(p, g, entries)
	case core.InConcurrentMode:
		err = doInConcurrent(p, g, entries)
	}
	return err
}

func doInSerial(p core.Parser, g core.Generator, entries []interface{}) error {
	mirTags, err := p.Parse(entries)
	if err == nil {
		return g.Generate(mirTags)
	}
	return err
}

func doInConcurrent(p core.Parser, g core.Generator, entries []interface{}) error {
	ctx := core.NewMirCtx(10)

	go p.GoParse(ctx, entries)
	go g.GoGenerate(ctx)

	select {
	case <-ctx.Done():
		if ctx.IsGeneratorDone() {
			return nil
		}
		return ctx.Err()
	}
}
