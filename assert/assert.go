// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assert

import (
	"github.com/alimy/mir/v4"
)

var (
	_typeAssertor TypeAssertor = anyTypeAssertor[fakeType]{}
)

// fakeType just a fake type for default type assertor
type fakeType struct{}

// Binding[T] binding interface for custom T context
type Binding[T any] interface {
	Bind(T) mir.Error
}

// Render[T] render interface for custom T context
type Render[T any] interface {
	Render(T)
}

// TypeAssertor type assert for Binding and Render interface
type TypeAssertor interface {
	AssertBinding(any) bool
	AssertRender(any) bool
}

// Register register custom TypeAssertor to assert Binding/Render interface
func Register(ta TypeAssertor) {
	_typeAssertor = ta
}

// RegisterType register custom TypeAssertor to assert Binding[T]/Render[T] interface
func RegisterType[T any]() {
	_typeAssertor = anyTypeAssertor[T]{}
}

// AssertBinding assert Binding interface for obj
func AssertBinding(obj any) bool {
	return _typeAssertor.AssertBinding(obj)
}

// AssertRender assert Render interface for obj
func AssertRender(obj any) bool {
	return _typeAssertor.AssertRender(obj)
}
