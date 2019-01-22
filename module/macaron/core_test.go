// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package macaron_test

import (
	"github.com/alimy/mir"
	"github.com/go-macaron/macaron"
	"net/http"
	"strings"
)

type entry struct {
	count uint64

	Chain    mir.Chain `mir:"-"`
	Group    mir.Group `mir:"v1"`
	add      mir.Post  `mir:"/add/:id/"`
	index    mir.Any   `mir:"/index/"`
	articles mir.Get   `mir:"/articles/:category/#GetArticles"`
}

// Add handler of "/add/:id"
func (e *entry) Add(c *macaron.Context) {
	body, _ := c.Req.Body().String()
	result := strings.Join([]string{
		"Add",
		c.Params("id"),
		body,
	}, ":")
	c.Resp.WriteHeader(http.StatusOK)
	c.Resp.Header().Add("Content-Type", "text/plain; charset=utf-8")
	c.Resp.Write([]byte(result))
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (e *entry) Index(c *macaron.Context) {
	e.count++
	c.Resp.WriteHeader(http.StatusOK)
	c.Resp.Header().Add("Content-Type", "text/plain; charset=utf-8")
	c.Resp.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (e *entry) GetArticles(c *macaron.Context) {
	c.Resp.WriteHeader(http.StatusOK)
	c.Resp.Header().Add("Content-Type", "text/plain; charset=utf-8")
	c.Resp.Write([]byte("GetArticles:" + c.Params("category")))
}

// mirChain chain used to register to engine
func mirChain() []macaron.Handler {
	return []macaron.Handler{
		macaron.Logger(),
		macaron.Recovery(),
	}
}
