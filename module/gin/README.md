# Mir.Gin
Mir.Gin module provider mir.Engine implement backend [Gin](https://github.com/gin-gonic/gin).

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
	v1 mir.Group `mir:"v1"`
	index mir.Get `mir:"/index/"`
	articles mir.Get `mir:"/articles/:category/#GetArticles"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(context gin.Context) {
	context.String(http.StatusOK, "get index data")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/:category/
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (h *site) GetArticles(context gin.Context) {
	context.String(http.StatusOK, "get articles data")
}

func main() {
	engine := gin.Default()         // Default gin engine
	mirE := ginE.Mir(engine)        // instance a mir engine
	mir.Register(mirE, &site{})     // Register handler to engine by mir
	engine.Run()                    // Start gin engine serve
}

```