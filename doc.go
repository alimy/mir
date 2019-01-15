// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

/*
Package mir provider yet another style to write http handler used register to router.

Define handler in struct type like below:

type site struct {
	Chain mir.Chain     `mir:"-"`
	Group mir.Group     `mir:"v1"`
	index mir.Get       `mir:"/index/"`
	articles mir.Get    `mir:"//localhost:8013/articles/{category}/{id:[0-9]+}?filter={filter}&foo=bar&id={id:[0-9]+}#GetArticles"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(c gin.Context) {
	c.String(http.StatusOK, "get index data")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Host info is the first segment start with '//'(eg:{subdomain}.domain.com)
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/{category}/{id:[0-9]+}?{filter})
// Queries info is the third info start with '?' and delimiter by '&'(eg: {filter}&{pages})
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles)
func (h *site) GetArticles(c gin.Context) {
	c.String(http.StatusOK, "get articles data")
}

Then register entry such use gin engine:

import (
	"github.com/alimy/mir"
	"github.com/gorilla/mux"
	"net/http"
	"log"

	muxE "github.com/alimy/mir/module/mux"
)

func main() {
	// Create a new mux router instance
	r := mux.NewRouter()

	// Register handler to engine by mir
	mir.SetDefault(muxE.Mir(r))
	mir.Register(&site{})

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
*/

package mir
