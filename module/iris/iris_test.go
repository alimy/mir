// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package iris

import (
	"github.com/alimy/mir"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"testing"
)

type site struct {
	count    uint64
	v1       mir.Group `mir:"v1"`
	add      mir.Post  `mir:"/add/{id:uint64}"`
	index    mir.Get   `mir:"/index/"`
	articles mir.Get   `mir:"/articles/{category:string}/#GetArticles"`
}

type blog struct {
	Chain    mir.Chain `mir:"-"`
	Group    mir.Group `mir:"v1"`
	articles mir.Get   `mir:"/articles/{name:string range(1,200) else 400}"`
}

// Add handler of "/add/:id"
func (h *site) Add(c context.Context) {
	c.Write([]byte("Add"))
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(c context.Context) {
	h.count++
	c.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (h *site) GetArticles(c context.Context) {
	c.Write([]byte("GetArticles"))
}

// Articles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (b *blog) Articles(c context.Context) {
	c.Write([]byte("Articles"))
}

func TestMir(t *testing.T) {
	app := iris.New()
	mirE := Mir(app)
	entries := mirEntries()
	if err := mir.Register(mirE, entries...); err != nil {
		t.Error(err)
	}
	// TODO: add httptest assert
}

// get all entries to register
func mirEntries() []interface{} {
	return []interface{}{
		&site{},
		&blog{
			Group: "v2", // direct custom group to v2 override default v1 in mir tag defined
			Chain: context.Handlers{
				recover.New(),
				logger.New(),
			},
		},
	}
}
