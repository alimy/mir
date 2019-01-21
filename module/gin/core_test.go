// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package gin_test

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
	"io"
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
func (e *entry) Add(c *gin.Context) {
	body, err := e.bytesFromBody(c.Request)
	if err != nil {
		body = []byte("")
	}
	result := strings.Join([]string{
		"Add",
		c.Params.ByName("id"),
		string(body),
	}, ":")
	c.String(http.StatusOK, result)
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (e *entry) Index(c *gin.Context) {
	e.count++
	c.String(http.StatusOK, "Index")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (e *entry) GetArticles(c *gin.Context) {
	c.String(http.StatusOK, "GetArticles:"+c.Params.ByName("category"))
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
func mirChain() gin.HandlersChain {
	return gin.HandlersChain{
		gin.Logger(),
		gin.Recovery(),
	}
}
