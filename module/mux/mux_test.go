// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mux

import (
	"github.com/alimy/mir"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

type site struct {
	count    uint32
	v1       mir.Group `mir:"v1"`
	add      mir.Post  `mir:"/add/{id}"`
	index    mir.Get   `mir:"/index/"`
	articles mir.Get   `mir:"//localhost:8013/articles/{category}/{id:[0-9]+}?filter={filter}&foo=bar&id={id:[0-9]+}#GetArticles"`
}

// Add handler of "/add/:id"
func (h *site) Add(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Add"))
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(rw http.ResponseWriter, r *http.Request) {
	h.count++
	rw.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (h *site) GetArticles(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GetArticles"))
}

func TestMir(t *testing.T) {
	r := mux.NewRouter()
	mirE := Mir(r)
	if err := mir.Register(mirE, &site{}); err != nil {
		t.Error(err)
	}
	// TODO: add httptest assert
}

func TestRegister(t *testing.T) {
	r := mux.NewRouter()
	if err := Register(r, &site{}); err != nil {
		t.Error(err)
	}
	// TODO: add httptest assert
}
