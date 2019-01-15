# Mir [![Build Status](https://api.travis-ci.com/alimy/mir.svg?branch=master)](https://travis-ci.com/alimy/mir)
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
	Chain mir.Chain     `mir:"-"`
	Group mir.Group     `mir:"v1"`
	index mir.Get       `mir:"/index/"`
	articles mir.Get    `mir:"/articles/:category/#GetArticles"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(c gin.Context) {
	c.String(http.StatusOK, "get index data")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/:category/#GetArticles)
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (h *site) GetArticles(c gin.Context) {
	c.String(http.StatusOK, "get articles data")
}

func main() {
	//Create a new gin engine
	engine := gin.New()
	
	// Register handler to engine by mir
	mirE := ginE.Mir(engine)
	mir.Register(mirE, &site{Chain: gin.HandlersChain{gin.Logger()}})
	
	// Start gin engine serve
	engine.Run()
}

```