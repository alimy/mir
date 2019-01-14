// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package gin

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type site struct {
	count    uint64
	v1       mir.Group `mir:"v1"`
	add      mir.Post  `mir:"/add/:id"`
	index    mir.Get   `mir:"/index/"`
	articles mir.Get   `mir:"/articles/:category/#GetArticles"`
}

// Add handler of "/add/:id"
func (h *site) Add(context *gin.Context) {
	context.String(http.StatusOK, "add")
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(context *gin.Context) {
	h.count++
	context.String(http.StatusOK, "Index")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (h *site) GetArticles(context *gin.Context) {
	context.String(http.StatusOK, "GetArticles")
}

func TestMir(t *testing.T) {
	engine := gin.Default()
	mir.SetDefault(Mir(engine))
	if err := mir.RegisterDefault(&site{}); err != nil {
		t.Error(err)
	}
	// TODO: add httptest assert
}
