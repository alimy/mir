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
	return p.reflex(entries)
}

// GoParse concurrent parse interface defined object entries
func (p *mirParser) GoParse(ctx core.MirCtx, entries []interface{}) {
	// start an *core.IfaceDescriptor deliver goroutine
	_, sink := ctx.Pipe()
	ifaceChan := make(chan *core.IfaceDescriptor, len(sink))
	go p.ifaceDeliver(ctx, ifaceChan)

	// concurrent parse entries
	var (
		err   error
		iface *core.IfaceDescriptor
	)

	wg := &sync.WaitGroup{}
	for _, entry := range entries {
		select {
		case <-ctx.Done():
			return
		default:
			go func() {
				defer func() {
					recover() // avoid send on closed channel error
				}()

				wg.Add(1)
				defer wg.Done()

				iface, err = p.ifaceFrom(entry)
				if err != nil {
					ctx.Cancel(err)
				}
				ifaceChan <- iface
			}()
		}
	}
	wg.Wait()

	// parse entries done and mark finish status
	close(ifaceChan)
}

// Clone return a copy of Parser
func (p *mirParser) Clone() core.Parser {
	return &mirParser{
		tagName: p.tagName,
	}
}
