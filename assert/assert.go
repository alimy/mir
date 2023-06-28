// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assert

var (
	_typeAssertor TypeAssertor = fakeTypeAssertor{}
)

// TypeAssertor type assert for Binding and Render interface
type TypeAssertor interface {
	AssertBinding(any) bool
	AssertRender(any) bool
}

// Register register custom TypeAssertor to assert Binding/Render interface
func Register(ta TypeAssertor) {
	_typeAssertor = ta
}

// AssertBinding assert Binding interface for obj
func AssertBinding(obj any) bool {
	return _typeAssertor.AssertBinding(obj)
}

// AssertRender assert Render interface for obj
func AssertRender(obj any) bool {
	return _typeAssertor.AssertRender(obj)
}
