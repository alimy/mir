// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"errors"
	"sync"

	"github.com/alimy/mir/v2/core"
	"github.com/alimy/mir/v2/internal/container"
)

func init() {
	core.RegisterParsers(&mirParser{tagName: defaultTag})
}

// mirParser parse for struct tag
type mirParser struct {
	tagName string
}

// Name name of parser
func (p *mirParser) Name() string {
	return core.ParserStructTag
}

// Init init parser
func (p *mirParser) Init(opts *core.ParserOpts) error {
	if opts == nil {
		return errors.New("init opts is nil")
	}
	p.tagName = opts.DefaultTag
	if p.tagName == "" {
		p.tagName = defaultTag
	}
	return nil
}

// Parse serial parse interface defined object entries
func (p *mirParser) Parse(entries []interface{}) (core.Descriptors, error) {
	if len(entries) == 0 {
		return nil, errors.New("entries is empty")
	}
	r := &reflex{tagName: p.tagName}
	return r.parse(entries)
}

// ParseContext concurrent parse interface defined object entries
func (p *mirParser) ParseContext(ctx core.MirCtx, entries []interface{}) {
	_, ifaceSink := ctx.Pipe()
	muxSet := container.NewMuxSet()

	wg := &sync.WaitGroup{}
	for _, entry := range entries {
		wg.Add(1)
		go func(ctx core.MirCtx, wg *sync.WaitGroup, ifaceSink chan<- *core.IfaceDescriptor, tagName string, entry interface{}) {
			defer wg.Done()

			r := &reflex{tagName: tagName}
			iface, err := r.ifaceFrom(entry)
			if err != nil {
				ctx.Cancel(err)
				return
			}
			core.Logus("parsed iface: %s.%s", iface.PkgName, iface.TypeName)
			if err = muxSet.Add(iface.PkgName + iface.TypeName); err != nil {
				ctx.Cancel(err)
				return
			}
			ifaceSink <- iface
			core.Logus("delivered iface: %s.%s", iface.PkgName, iface.TypeName)
		}(ctx, wg, ifaceSink, p.tagName, entry)
	}
	wg.Wait()

	ctx.ParserDone()
}

// Clone return a copy of Parser
func (p *mirParser) Clone() core.Parser {
	return &mirParser{
		tagName: p.tagName,
	}
}
