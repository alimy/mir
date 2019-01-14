// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
	"net/http"

	ginE "github.com/alimy/mir/module/gin"
)

type site struct {
	v1       mir.Group `mir:"v1"`
	index    mir.Get   `mir:"/index/"`
	articles mir.Get   `mir:"/articles/:category/#GetArticles"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(context *gin.Context) {
	context.String(http.StatusOK, "get index data")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (h *site) GetArticles(context *gin.Context) {
	context.String(http.StatusOK, "get articles data")
}

func main() {
	engine := gin.Default()

	// Register handler to engine by mir
	e := ginE.Mir(engine)
	if err := mir.Register(e, &site{}); err != nil {
		panic(err)
	}

	// Start gin engine serve
	engine.Run()
}
