// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package httprouter_test

import (
	"github.com/alimy/mir"
	"github.com/julienschmidt/httprouter"
	"io"
	"strings"

	"net/http"
)

type entry struct {
	count uint64

	add      mir.Post `mir:"/add/:id/"`
	index    mir.Get  `mir:"/index/"`
	articles mir.Get  `mir:"/articles/:category/#GetArticles"`
}

// Add handler of "/add/:id"
func (e *entry) Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := e.bytesFromBody(r)
	if err != nil {
		body = []byte("")
	}
	result := strings.Join([]string{
		"Add",
		p.ByName("id"),
		string(body),
	}, ":")
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (e *entry) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (e *entry) GetArticles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("GetArticles:" + p.ByName("category")))
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
