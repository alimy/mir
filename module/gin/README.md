# Mir.Gin
Mir.Gin module provider mir.Engine implement backend [Gin](https://github.com/gin-gonic/gin).

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

func (h *site) Index(context gin.Context) {
	context.String(http.StatusOK, "get index data")
}

func main() {
	// Default gin engine
	engine := gin.Default()

    // Setup mir engine
	mir.Setup(ginE.Mir(egine))
	
	// Register handler to engine by mir
	mir.Register(&site{})
	
	// Start serve
	engine.Run()
}

```