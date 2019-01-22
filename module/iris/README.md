# Mir.Iris
Mir.Iris module provide mir.Engine implement backend [Iris](https://github.com/kataras/iris).

### Usage 
```go
package main

import(
	"github.com/alimy/mir"
    "github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"net/http"

	mirE "github.com/alimy/mir/module/iris"
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
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/:category)
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (h *site) GetArticles(c context.Context) {
	c.Write([]byte("GetArticles"))
}

// Articles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (b *blog) Articles(c context.Context) {
    c.Write([]byte("Articles"))
}

func main() {
	// Create a new iris.Application instance
	app := iris.New()          
	
	// Register handler to engine by mir
	entries := mirEntries()
	mirE.Register(app, entries...)
	
	// Start iris application
	app.Run(iris.Addr(":8013"))
}

/// get all entries to register
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
```