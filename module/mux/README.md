# Mir.Mux
Mir.Mux module provider mir.Engine implement backend [Mux](https://github.com/gorilla/mux).

### Usage 
```go
package main

import(
	"github.com/alimy/mir"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	
	mirE "github.com/alimy/mir/module/mux"
)

type site struct {
	count    uint32
	Group mir.Group     `mir:"v1"`
	index mir.Get       `mir:"/index/"`
	articles mir.Get    `mir:"//localhost:8013/articles/{category}/{id:[0-9]+}?filter={filter}&foo=bar&id={id:[0-9]+}#GetArticles"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(rw http.ResponseWriter, r *http.Request) {
	h.count++
	rw.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Host info is the first segment start with '//'(eg:{subdomain}.domain.com)
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/{category}/{id:[0-9]+}?{filter})
// Queries info is the third info start with '?' and delimiter by '&'(eg: {filter}&{pages})
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (h *site) GetArticles(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GetArticles"))
}

func main() {
	// Create a new mux router instance
	r := mux.NewRouter()
	
	// Instance a mir engine to register handler for mux router by mir
	mirE.Register(r, &site{}) 
	
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r)) 
}

```