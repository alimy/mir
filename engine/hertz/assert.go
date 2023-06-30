// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine_hertz

import (
	"github.com/alimy/mir/v4"
	"github.com/alimy/mir/v4/assert"
	"github.com/cloudwego/hertz/pkg/app"
)

func init() {
	assert.Register(typeAssortor{})
}

type Binding interface {
	Bind(*app.RequestContext) mir.Error
}

type Render interface {
	Render(*app.RequestContext)
}

type typeAssortor struct{}

func (typeAssortor) AssertBinding(obj any) bool {
	_, ok := obj.(Binding)
	return ok
}

func (typeAssortor) AssertRender(obj any) bool {
	_, ok := obj.(Render)
	return ok
}
