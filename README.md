# Mir
Mir is a tookit for register handler to http engine router(eg: [Gin](https://github.com/gin-gonic/gin),[echo](https://github.com/labstack/echo),[mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter))
 use struct tag info.

### Usage 
```go
package main

import(
	"github.com/gin-gonic/gin"
	"github.com/alimy/mir"
	"net/http"
	
	ginE "github.com/alimy/mir/module/gin"
)

type site struct {
	Chain mir.Chain `mir:"-"`
	Group mir.Group `mir:"v1"`
	index mir.Get `mir:"/index/"`
	articles mir.Get `mir:"//{subdomain}.domain.com/articles/{category}/{id:[0-9]+}?{filter}&{pages}#GetArticles"`
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
// Host info is the first segment start with '//'(eg:{subdomain}.domain.com)
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/{category}/{id:[0-9]+}?{filter})
// Queries info is the third info start with '?' and delimiter by '&'(eg: {filter}&{pages})
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
	engine := gin.New()             // Default gin engine
	
	mirE := ginE.Mir(engine)        // instance a mir engine
	entries := mirEntries()
	mir.Register(mirE, entries...)  // Register handler to engine by mir
	
	engine.Run()                    // Start gin engine serve
}

// get all entries to register
func mirEntries()[]interface{} {
	return []interface{} {
		&site{},
		&blog{
			Group:"v2", // direct custom group to v2 override default v1 in mir tag defined
			Chain: gin.HandlersChain {
				gin.Logger(),
	            gin.Recovery(),
			}},
	}
}

```