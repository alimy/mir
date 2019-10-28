// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package gin_test

import (
	"bytes"

	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
	"net/http/httptest"

	. "github.com/alimy/mir/module/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {
	var (
		engine *gin.Engine
		w      *httptest.ResponseRecorder
		err    error
	)

	JustBeforeEach(func() {
		w = httptest.NewRecorder()

	})

	Context("check Mir function", func() {
		BeforeEach(func() {
			engine = gin.New()
			mirE := Mir(engine)
			err = mir.Register(mirE, &entry{Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(engine).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v1/add/10086/", body)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/index/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/articles/golang/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})
	})

	Context("check Register function", func() {
		BeforeEach(func() {
			engine = gin.New()
			err = Register(engine, &entry{Group: "v2", Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(engine).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})
	})

	Context("check Register entries", func() {
		BeforeEach(func() {
			engine = gin.New()
			err = Register(engine, &entry{}, &entry{Group: "v2", Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(engine).ShouldNot(BeNil())
		})

		It("handle v1 add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v1/add/10086/", body)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler v1 index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/index/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle v1 articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/articles/golang/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})

		It("handle v2 add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler v2 index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle v2 articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/", nil)
			engine.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang"))
		})
	})
})
