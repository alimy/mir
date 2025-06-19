// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"errors"
	"sync"

	"github.com/alimy/mir/v5/core"
	"github.com/alimy/mir/v5/internal/reflex"
	"github.com/alimy/mir/v5/internal/utils"
)

var (
	// defaultTag indicate default mir's struct tag name
	defaultTag        = "mir"
	defaultBindingTag = "binding"
	defaultRenderTag  = "render"
	defautlMethodTag  = "method"
)

func init() {
	core.RegisterParsers(&mirParser{
		engineInfo: &core.EngineInfo{},
		tagName:    defaultTag,
	})
}

// mirParser parse for struct tag
type mirParser struct {
	engineInfo      *core.EngineInfo
	tagName         string
	bindingTagName  string
	renderTagName   string
	watchCtxDone    bool
	useRequestCtx   bool
	useNamedBinding bool
	useNamedRender  bool
	noneQuery       bool
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
	if opts.EngineInfo != nil {
		p.engineInfo = opts.EngineInfo
	}
	p.tagName = opts.DefaultTag
	p.bindingTagName = opts.DefaultBindingTag
	p.renderTagName = opts.DefaultRenderTag
	p.watchCtxDone = opts.WatchCtxDone
	p.useRequestCtx = opts.UseRequestCtx
	p.useNamedBinding = opts.UseNamedBinding
	p.useNamedRender = opts.UseNamedRender
	p.noneQuery = opts.NoneQuery
	if p.tagName == "" {
		p.tagName = defaultTag
	}
	if p.bindingTagName == "" {
		p.bindingTagName = defaultBindingTag
	}
	if p.renderTagName == "" {
		p.renderTagName = defaultRenderTag
	}
	return nil
}

// Parse serial parse interface defined object entries
func (p *mirParser) Parse(entries []any) (core.Descriptors, error) {
	if len(entries) == 0 {
		return nil, errors.New("entries is empty")
	}
	r := reflex.NewReflex(p.engineInfo,
		p.tagName, p.bindingTagName, p.renderTagName,
		p.useNamedBinding, p.useNamedRender, p.watchCtxDone, p.useRequestCtx, p.noneQuery)
	return r.Parse(entries)
}

// ParseContext concurrent parse interface defined object entries
func (p *mirParser) ParseContext(ctx core.MirCtx, entries []any) {
	_, ifaceSink := ctx.Pipe()
	muxSet := utils.NewMuxSet()

	wg := &sync.WaitGroup{}
	for _, entry := range entries {
		wg.Add(1)
		go func(ctx core.MirCtx, wg *sync.WaitGroup, ifaceSink chan<- *core.IfaceDescriptor, entry any) {
			defer wg.Done()

			r := reflex.NewReflex(p.engineInfo,
				p.tagName, p.bindingTagName, p.renderTagName,
				p.useNamedBinding, p.useNamedRender, p.watchCtxDone, p.useRequestCtx, p.noneQuery)
			iface, err := r.IfaceFrom(entry)
			if err != nil {
				core.Logus("ifaceFrom error: %s", err)
				ctx.Cancel(err)
				return
			}
			// no actual fields so just continue
			if len(iface.Fields) == 0 {
				return
			}
			core.Logus("parsed iface: %s.%s", iface.PkgName, iface.TypeName)
			if err = muxSet.Add(iface.Group + iface.TypeName); err != nil {
				core.Logus("muxSet.Add error: %s", err)
				ctx.Cancel(err)
				return
			}
			ifaceSink <- iface
			core.Logus("delivered iface: %s.%s", iface.PkgName, iface.TypeName)
		}(ctx, wg, ifaceSink, entry)
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
