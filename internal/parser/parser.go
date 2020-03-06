// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"errors"

	"github.com/alimy/mir/v2/core"
)

func init() {
	core.RegisterParsers(parserStructTag{})
}

// parserStructTag parse for struct tag
type parserStructTag struct{}

// Name name of parser
func (parserStructTag) Name() string {
	return core.ParserStructTag
}

// Parse parse interface define object entries
func (parserStructTag) Parse(entries []interface{}) (core.Descriptors, error) {
	if len(entries) == 0 {
		return nil, errors.New("entries is empty")
	}
	return reflex(entries)
}

// SetDefaultTag set default tag name
func SetDefaultTag(tag string) {
	if len(tag) > 0 {
		defaultTag = tag
	}
}

// DefaultTag return default tag name
func DefaultTag() string {
	return defaultTag
}
