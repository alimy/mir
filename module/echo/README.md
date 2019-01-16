# Mir.Echo
Mir.Echo module provider mir.Engine implement backend [Echo](https://github.com/labstack/echo).

### Usage 
```go
package main

import(
	"github.com/alimy/mir"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

	
	mirE "github.com/alimy/mir/module/echo"
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
func (h *site) Index(c echo.Context) error {
	return c.String(http.StatusOK, "get index data")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/:category)
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (h *site) GetArticles(c echo.Context) error {
	return c.String(http.StatusOK, "get articles data")
}

// Articles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (b *blog) Articles(c echo.Context) error {
	return c.String(http.StatusOK, "get articles data")
}

func main() {
	// Create a new echo engine instance
	e := echo.New()             
	
	// Register handler to engine by mir
	entries := mirEntries()
	mirE.Register(e, entries...)
	
	// Start echo engine serve
	e.Start(":8013")
}

// get all entries to register
func mirEntries()[]interface{} {
	return []interface{} {
		&site{},
		&blog{
			Group: "v2", // direct custom group to v2 override default v1 in mir tag defined
			Chain: []echo.MiddlewareFunc {
				middleware.Logger(),
	            middleware.Recover(),
			}},
	}
}
```