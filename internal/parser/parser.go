// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"errors"
	"sync"

	"github.com/alimy/mir/v2/core"
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
func (p *mirParser) Init(opts core.InitOpts) error {
	if len(opts) != 0 {
		p.tagName = opts[core.OptDefaultTag]
	}
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

// GoParse concurrent parse interface defined object entries
func (p *mirParser) GoParse(ctx core.MirCtx, entries []interface{}) {
	// start an *core.IfaceDescriptor deliver goroutine
	ifaceChan := make(chan *core.IfaceDescriptor, ctx.ChanCapcity())
	parserDone := make(chan struct{})
	go p.ifaceDeliver(ctx, ifaceChan, parserDone)

	// concurrent parse entries
	var (
		err   error
		iface *core.IfaceDescriptor
	)

	wg := &sync.WaitGroup{}
	wg.Add(len(entries))
	for _, entry := range entries {
		select {
		case <-ctx.Done():
			return
		default:
			go func(ifaceSink chan<- *core.IfaceDescriptor, tagName string, entry interface{}) {
				wg.Done()
				r := &reflex{tagName: tagName}
				iface, err = r.ifaceFrom(entry)
				if err != nil {
					ctx.Cancel(err)
				}
				ifaceSink <- iface
				core.Logus("deliver iface: %s.%s", iface.PkgName, iface.TypeName)
			}(ifaceChan, p.tagName, entry)
		}
	}
	wg.Wait()

	// parse entries done and mark finish status
	close(parserDone)
}

func (p *mirParser) ifaceDeliver(ctx core.MirCtx, source <-chan *core.IfaceDescriptor, parserDone <-chan struct{}) {
	var err error
	_, ifaceSink := ctx.Pipe()
	ds := make(core.Descriptors, ctx.ChanCapcity())
	for {
		select {
		case iface := <-source:
			if len(iface.Fields) == 0 {
				continue
			}
			if err = ds.Put(iface); err != nil {
				ctx.Cancel(err)
				return
			}
			ifaceSink <- iface
		case <-ctx.Done():
			return
		case <-parserDone:
			ctx.ParserDone()
			return
		}
	}
}

// Clone return a copy of Parser
func (p *mirParser) Clone() core.Parser {
	return &mirParser{
		tagName: p.tagName,
	}
}
