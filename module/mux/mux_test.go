// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mux_test

import (
	"bytes"
	"github.com/alimy/mir"
	"github.com/gorilla/mux"
	"net/http/httptest"

	. "github.com/alimy/mir/module/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {
	var (
		router *mux.Router
		w      *httptest.ResponseRecorder
		err    error
	)

	JustBeforeEach(func() {
		w = httptest.NewRecorder()
	})

	Context("check Mir function", func() {
		BeforeEach(func() {
			router = mux.NewRouter()
			mirE := Mir(router)
			err = mir.Register(mirE, &entry{Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v1/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/articles/golang/10086?filter=module&foo=bar&num=5", nil)
			r.Host = "alimy.example.com"
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:alimy:golang:10086:module:5"))
		})
	})

	Context("check Register function", func() {
		BeforeEach(func() {
			router = mux.NewRouter()
			err = Register(router, &entry{Group: "/v2", Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/10086?filter=module&foo=bar&num=5", nil)
			r.Host = "alimy.example.com"
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:alimy:golang:10086:module:5"))
		})
	})

	Context("check Register entries", func() {
		BeforeEach(func() {
			router = mux.NewRouter()
			err = Register(router, &entry{}, &entry{Group: "v2", Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle v1 add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler v1 index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle v1 articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/10086?filter=module&foo=bar&num=5", nil)
			r.Host = "alimy.example.com"
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:alimy:golang:10086:module:5"))
		})

		It("handle v2 add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler v2 index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle v2 articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/10086?filter=module&foo=bar&num=5", nil)
			r.Host = "alimy.example.com"
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:alimy:golang:10086:module:5"))
		})
	})
})
