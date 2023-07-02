// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine_mux

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/alimy/mir/v4/assert"
)

func init() {
	assert.Register(typeAssertor{})
}

type Binding interface {
	Bind(*http.Request) mir.Error
}

type Render interface {
	Render(http.ResponseWriter)
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
