// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine_gin

import (
	"github.com/alimy/mir/v4"
	"github.com/alimy/mir/v4/assert"
	"github.com/gin-gonic/gin"
)

func init() {
	assert.Register(typeAssertor{})
}

type Binding interface {
	Bind(*gin.Context) mir.Error
}

type Render interface {
	Render(*gin.Context)
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
