// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine_iris

import (
	"github.com/alimy/mir/v4"
	"github.com/alimy/mir/v4/assert"
	"github.com/kataras/iris/v12/context"
)

func init() {
	assert.Register(typeAssertor{})
}

type Binding interface {
	Bind(*context.Context) mir.Error
}

type Render interface {
	Render(*context.Context)
}

type typeAssertor struct{}

func (typeAssertor) AssertBinding(obj any) bool {
	_, ok := obj.(Binding)
	return ok
}

func (typeAssertor) AssertRender(obj any) bool {
	_, ok := obj.(Render)
	return ok
}
