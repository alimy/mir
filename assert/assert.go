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

// Binding2[R, P] binding interface for custom T context
type Binding2[R, P any] interface {
	Bind(R, P) mir.Error
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

// RegisterType[T] register custom TypeAssertor to assert Binding[T]/Render[T] interface
func RegisterType[T any]() {
	_typeAssertor = anyTypeAssertor[T]{}
}

// RegisterType2[B, R] register custom TypeAssertor to assert Binding[B]/Render[R] interface
func RegisterType2[B, R any]() {
	_typeAssertor = anyTypeAssertor2[B, R]{}
}

// RegisterType2[B, P, R] register custom TypeAssertor to assert Binding[B]/Render[R] interface
func RegisterType3[B, P, R any]() {
	_typeAssertor = anyTypeAssertor3[B, P, R]{}
}

// AssertBinding assert Binding interface for obj
func AssertBinding(obj any) bool {
	return _typeAssertor.AssertBinding(obj)
}

// AssertRender assert Render interface for obj
func AssertRender(obj any) bool {
	return _typeAssertor.AssertRender(obj)
}
