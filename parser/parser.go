// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"errors"

	"github.com/alimy/mir/v2/core"
)

func init() {
	core.SetDefParser(parserStructTag{})
}

// parserStructTag parse for struct tag
type parserStructTag struct{}

// Name name of parser
func (parserStructTag) Name() string {
	return "parserStructTag"
}

// Parse parse interface define object entries
func (parserStructTag) Parse(entries []interface{}) ([]*core.TagMir, error) {
	// TODO
	return nil, errors.New("not ready")
}
