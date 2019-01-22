# Mir.Macaron
Mir.Macaron module provide mir.Engine implement backend by [Macaron](https://github.com/go-macaron/macaron).

### Usage 
* Get module first
```bash
$ go get github.com/alimy/mir/module/macaron@master
```

* Then happy in your heart to codding...

```go
package main

import(
	"github.com/alimy/mir"
	"github.com/go-macaron/macaron"
	"net/http"
	
	mirE "github.com/alimy/mir/module/macaron"
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
func (s *site) Index(c *macaron.Context) {
	c.Resp.WriteHeader(http.StatusOK)
	c.Resp.Header().Add("Content-Type", "text/plain; charset=utf-8")
	c.Resp.Write([]byte("Index"))
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/:category)
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (s *site) GetArticles(c *macaron.Context) {
	c.Resp.WriteHeader(http.StatusOK)
	c.Resp.Header().Add("Content-Type", "text/plain; charset=utf-8")
	c.Resp.Write([]byte("GetArticles"))
}

// Articles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (b *blog) Articles(c *macaron.Context) {
	c.Resp.WriteHeader(http.StatusOK)
	c.Resp.Header().Add("Content-Type", "text/plain; charset=utf-8")
	c.Resp.Write([]byte("Articles"))
}

func main() {
	// Default gin engine
	m := macaron.New()             
	
	// Register handler to engine by mir
	entries := mirEntries()
	mirE.Register(m, entries...)
	
	 // Start gin engine serve
	m.Run()                   
}

// get all entries to register
func mirEntries()[]interface{} {
	return []interface{} {
		&site{},
		&blog{
			Group: "v2", // direct custom group to v2 override default v1 in mir tag defined
			Chain: []macaron.Handler {
				macaron.Logger(),
	            		macaron.Recovery(),
			},
		},
	}
}
```
