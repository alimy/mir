// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package iris_test

import (
	"testing"

	"github.com/alimy/mir"
	"github.com/gavv/httpexpect"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"

	. "github.com/alimy/mir/module/iris"
)

func TestRegister(t *testing.T) {
	app := iris.New()
	Register(app, &entry{Chain: mirChain()}, &entry{Group: "v2"})
	expect := httptest.New(t, app, httptest.Debug(false))
	testEntry(expect)
}

func TestMir(t *testing.T) {
	app := iris.New()
	mirE := Mir(app)
	mir.Register(mirE, &entry{Chain: mirChain()}, &entry{Group: "v2"})
	expect := httptest.New(t, app, httptest.Debug(false))
	testEntry(expect)
}

func testEntry(expect *httpexpect.Expect) {
	for _, test := range []struct {
		method string
		path   string
		body   []byte
		status int
		result string
	}{
		{
			method: mir.MethodPost,
			path:   "/v1/add/10086",
			body:   []byte("hello"),
			status: httptest.StatusOK,
			result: "Add:10086:hello",
		},
		{
			method: mir.MethodGet,
			path:   "/v1/index",
			body:   []byte(""),
			status: httptest.StatusOK,
			result: "Index",
		},
		{
			method: mir.MethodGet,
			path:   "/v1/articles/golang",
			status: httptest.StatusOK,
			result: "GetArticles:golang",
		},
		{
			method: mir.MethodPost,
			path:   "/v2/add/10086",
			body:   []byte("hello"),
			status: httptest.StatusOK,
			result: "Add:10086:hello",
		},
		{
			method: mir.MethodGet,
			path:   "/v2/index",
			body:   []byte(""),
			status: httptest.StatusOK,
			result: "Index",
		},
		{
			method: mir.MethodGet,
			path:   "/v2/articles/golang",
			status: httptest.StatusOK,
			result: "GetArticles:golang",
		},
	} {
		expect.
			Request(test.method, test.path).
			WithBytes(test.body).
			Expect().
			Status(test.status).
			Text().Equal(test.result)
	}
}
