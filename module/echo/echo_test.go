// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package echo_test

import (
	"bytes"
	"github.com/alimy/mir"
	"github.com/labstack/echo"
	"net/http/httptest"

	. "github.com/alimy/mir/module/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {
	var (
		e   *echo.Echo
		w   *httptest.ResponseRecorder
		err error
	)

	JustBeforeEach(func() {
		w = httptest.NewRecorder()
	})

	Context("check Mir function", func() {
		BeforeEach(func() {
			e = echo.New()
			mirE := Mir(e)
			err = mir.Register(mirE, &entry{Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v1/add/10086/", body)
			e.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/index/", nil)
			e.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/articles/golang/", nil)
			e.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})
	})

	Context("check Register function", func() {
		BeforeEach(func() {
			e = echo.New()
			err = Register(e, &entry{Group: "v2", Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			e.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			e.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/", nil)
			e.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})
	})
})
