// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package chi_test

import (
	"io"
	"net/http"
	"strings"

	"github.com/alimy/mir"
	"github.com/go-chi/chi"
)

type entry struct {
	count uint32

	Chain      mir.Chain `mir:"-"`
	Group      mir.Group `mir:"v1"`
	add        mir.Post  `mir:"/add/{id}/"`
	index      mir.Any   `mir:"/index/"`
	articles   mir.Get   `mir:"/articles/{category}/{id:[0-9]+}#GetArticles"`
	chainFunc1 mir.Get   `mir:"/chainfunc1#-ChainFunc"`
	chainFunc2 mir.Get   `mir:"/chainfunc2#GetChainFunc2&ChainFunc"`
}

// Add handler of "/add/{id}"
func (e *entry) Add(rw http.ResponseWriter, r *http.Request) {
	body, err := e.bytesFromBody(r)
	if err != nil {
		body = []byte("")
	}
	result := strings.Join([]string{
		"Add",
		chi.URLParam(r, "id"),
		string(body),
	}, ":")
	rw.WriteHeader(200)
	rw.Write([]byte(result))
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (e *entry) Index(rw http.ResponseWriter, r *http.Request) {
	e.count++
	rw.WriteHeader(200)
	rw.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (e *entry) GetArticles(rw http.ResponseWriter, r *http.Request) {
	result := strings.Join([]string{
		"GetArticles",
		chi.URLParam(r, "category"),
		chi.URLParam(r, "id"),
	}, ":")
	rw.WriteHeader(200)
	rw.Write([]byte(result))
}

// ChainFunc1 handler with chain func info.
func (e *entry) ChainFunc1(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("ChainFunc1"))
}

// GetChainFunc2 handler with chain func info.
func (e *entry) GetChainFunc2(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("GetChainFunc2"))
}

// ChainFunc return field's online middleware
func (e *entry) ChainFunc() chi.Middlewares {
	return chi.Middlewares{
		simpleMiddleware,
		simpleMiddleware,
	}
}

// bytesFromBody get contents from request's body
func (e *entry) bytesFromBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()

	buf := [256]byte{}
	result := make([]byte, 0)
	if size, err := r.Body.Read(buf[:]); err == nil {
		result = append(result, buf[:size]...)
	} else if err != io.EOF {
		return nil, err
	}
	return result, nil
}

// mirChain chain used to register to engine
func mirChain() chi.Middlewares {
	return chi.Middlewares{
		simpleMiddleware,
		simpleMiddleware,
	}
}

func simpleMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do nothing just for test
		h.ServeHTTP(w, r)
	})
}
