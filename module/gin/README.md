# Mir.Gin
Mir.Gin module provider mir.Engine implement backend [Gin](https://github.com/gin-gonic/gin).

### Usage 
```go
package main

import(
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
	"net/http"
	
	mirE "github.com/alimy/mir/module/gin"
)

type site struct {
	Group mir.Group     `mir:"v1"`
	index mir.Get       `mir:"/index/"`
	articles mir.Get    `mir:"/articles/:category/#GetArticles"`
}

type blog struct {
	Chain mir.Chain     `mir:"-"`
	Group mir.Group     `mir:"v1"`
	articles mir.Get    `mir:"/articles/:category"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(c gin.Context) {
	c.String(http.StatusOK, "get index data")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/:category)
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (h *site) GetArticles(c gin.Context) {
	c.String(http.StatusOK, "get articles data")
}

// Articles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (b *blog) Articles(c gin.Context) {
	c.String(http.StatusOK, "get articles data")
}

func main() {
	// Default gin engine
	engine := gin.New()             
	
	// Register handler to engine by mir
	entries := mirEntries()
	mirE.Register(engine, entries...)
	
	 // Start gin engine serve
	engine.Run()                   
}

// get all entries to register
func mirEntries()[]interface{} {
	return []interface{} {
		&site{},
		&blog{
			Group: "v2", // direct custom group to v2 override default v1 in mir tag defined
			Chain: gin.HandlersChain {
				gin.Logger(),
	            gin.Recovery(),
			}},
	}
}
```