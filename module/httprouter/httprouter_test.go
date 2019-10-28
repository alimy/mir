// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package httprouter_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/alimy/mir"
	"github.com/julienschmidt/httprouter"

	. "github.com/alimy/mir/module/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {
	var (
		router *httprouter.Router
		w      *httptest.ResponseRecorder
		err    error
	)

	JustBeforeEach(func() {
		w = httptest.NewRecorder()

	})

	Context("check Mir function", func() {
		BeforeEach(func() {
			router = httprouter.New()
			mirE := Mir(router)
			err = mir.Register(mirE, &entry{})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/articles/golang/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})
	})

	Context("check Register function", func() {
		BeforeEach(func() {
			router = httprouter.New()
			err = Register(router, &entry{})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/articles/golang/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})
	})
})
