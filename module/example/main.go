// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
	"net/http"

	mirE "github.com/alimy/mir/module/gin"
)

type site struct {
	Chain    mir.Chain `mir:"-"`
	v1       mir.Group `mir:"v1"`
	index    mir.Get   `mir:"/index/"`
	articles mir.Get   `mir:"/articles/:category/#GetArticles"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (s *site) Index(c *gin.Context) {
	c.String(http.StatusOK, "Index")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (s *site) GetArticles(c *gin.Context) {
	c.String(http.StatusOK, "GetArticles:"+c.Param("category"))
}

func main() {
	e := gin.New()

	// Register handler to engine by mir
	handlersChain := gin.HandlersChain{gin.Logger(), gin.Recovery()}
	if err := mirE.Register(e, &site{Chain: handlersChain}); err != nil {
		panic(err)
	}

	// Start gin engine serve
	e.Run()
}
