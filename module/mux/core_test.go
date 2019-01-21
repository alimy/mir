// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mux_test

import (
	"github.com/alimy/mir"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strings"
)

type entry struct {
	count uint32

	Chain    mir.Chain `mir:"-"`
	Group    mir.Group `mir:"v1"`
	add      mir.Post  `mir:"/add/{id}/"`
	index    mir.Any   `mir:"/index/"`
	articles mir.Get   `mir:"//{subdomain:[a-z]+}.example.com/articles/{category}/{id:[0-9]+}?filter={filter}&foo=bar&num={num:[0-9]+}#GetArticles"`
}

// Add handler of "/add/{id}"
func (e *entry) Add(rw http.ResponseWriter, r *http.Request) {
	body, err := e.bytesFromBody(r)
	if err != nil {
		body = []byte("")
	}
	vars := mux.Vars(r)
	result := strings.Join([]string{
		"Add",
		vars["id"],
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
	vars := mux.Vars(r)
	result := strings.Join([]string{
		"GetArticles",
		vars["subdomain"],
		vars["category"],
		vars["id"],
		vars["filter"],
		vars["num"],
	}, ":")
	rw.WriteHeader(200)
	rw.Write([]byte(result))
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
func mirChain() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{
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
