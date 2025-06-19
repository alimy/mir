// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assert

import (
	"testing"
)

type (
	fakeCtx         struct{}
	fakeBindingCtx  struct{}
	fakeRenderCtx   struct{}
	fakeObjectType  struct{}
	fakeObjectType2 struct{}
)

func (fakeObjectType) Bind(_c fakeCtx) error {
	// nothing
	return nil
}

func (fakeObjectType) Render(_c fakeCtx) {
	// nothing
}

func (fakeObjectType2) Bind(_c fakeBindingCtx) error {
	// nothing
	return nil
}

func (fakeObjectType2) Render(_c fakeRenderCtx) {
	// nothing
}

func TestAssertType(t *testing.T) {
	var ta TypeAssertor = anyTypeAssertor[fakeCtx]{}
	fakeObj := new(fakeObjectType)
	if ok := ta.AssertBinding(fakeObj); !ok {
		t.Error("want assert binding true but not")
	}
	if ok := ta.AssertRender(fakeObj); !ok {
		t.Error("want assert render true but not")
	}
}

func TestAssertType2(t *testing.T) {
	var ta TypeAssertor = anyTypeAssertor2[fakeBindingCtx, fakeRenderCtx]{}
	fakeObj := new(fakeObjectType2)
	if ok := ta.AssertBinding(fakeObj); !ok {
		t.Error("want assert binding true but not")
	}
	if ok := ta.AssertRender(fakeObj); !ok {
		t.Error("want assert render true but not")
	}
}

func TestAssertBinding(t *testing.T) {
	fakeObj := new(fakeObjectType)
	if ok := AssertBinding(fakeObj); ok {
		t.Error("want assert binding false but not")
	}
	if ok := AssertRender(fakeObj); ok {
		t.Error("want assert render false but not")
	}
}

func TestRegisterType(t *testing.T) {
	fakeObj := new(fakeObjectType)
	RegisterType[fakeCtx]()
	if ok := AssertBinding(fakeObj); !ok {
		t.Error("want assert binding true but not")
	}
	if ok := AssertRender(fakeObj); !ok {
		t.Error("want assert render true but not")
	}
}

func TestRegisterType2(t *testing.T) {
	fakeObj := new(fakeObjectType2)
	RegisterType2[fakeBindingCtx, fakeRenderCtx]()
	if ok := AssertBinding(fakeObj); !ok {
		t.Error("want assert binding true but not")
	}
	if ok := AssertRender(fakeObj); !ok {
		t.Error("want assert render true but not")
	}
}
