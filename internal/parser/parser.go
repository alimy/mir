// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"errors"

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

// Parse parse interface define object entries
func (p *mirParser) Parse(entries []interface{}) (core.Descriptors, error) {
	if len(entries) == 0 {
		return nil, errors.New("entries is empty")
	}
	return p.reflex(entries)
}

// Clone return a copy of Parser
func (p *mirParser) Clone() core.Parser {
	return &mirParser{
		tagName: p.tagName,
	}
}
