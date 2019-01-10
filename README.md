# Mir
Mir is a tookit for register handler to http engine use struct tag info.

### Usage 
```go
package main

import(
	"github.com/gin-gonic/gin"
	"github.com/alimy/mir"
	"net/http"
	
	ginE "gihub.com/alimy/mir/module/gin"
)

type site struct {
	v1 mir.Group `mir:"v1"`
	index mir.Get `mir:"/index/"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(context gin.Context) {
	context.String(http.StatusOK, "get index data")
}

func main() {
	engine := gin.Default()     // Default gin engine
	mir.Setup(ginE.Mir(egine))  // Setup mir engine
	mir.Register(&site{})       // Register handler to engine by mir
	engine.Run()                // Start gin engine serve
}

```