// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

/*
Package mir provider yet another style to write http handler used register to router.

Define handler in struct type like below:

	type entry struct {
		count uint32

		Chain mir.Chain     `mir:"-"`
		Group mir.Group     `mir:"v1"`
		index mir.Get       `mir:"/index/"`
		articles mir.Get    `mir:"//{subdomain}.example.com/articles/{category}/{id:[0-9]+}?filter={filter}&foo=bar&num={num:[0-9]+}#GetArticles"`
	}

	// Index handler of the index field that in site struct, the struct tag indicate
	// this handler will register to path "/index/" and method is http.MethodGet.
	func (e *entry) Index(c Index(rw http.ResponseWriter, r *http.Request) {
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

Then register entry such use mux router:

	import (
		"github.com/alimy/mir"
		"github.com/gorilla/mux"
		"net/http"
		"log"

		mirE "github.com/alimy/mir/module/mux"
	)

	func main() {
		// Create a new mux router instance
		r := mux.NewRouter()

		// Register handler to engine by mir
		mirE.Register(r, &entry{Chain: mirChain()})

		// Bind to a port and pass our router in
		log.Fatal(http.ListenAndServe(":8000", r))
	}

*/
package mir
