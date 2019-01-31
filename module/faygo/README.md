# Mir.Faygo
Mir.Faygo module provide mir.Engine implement backend by [Faygo](https://github.com/henrylee2cn/faygo).

### Usage 
```go
package main

import (
	"github.com/alimy/mir"
	"github.com/henrylee2cn/faygo"
	"time"

	mirE "github.com/alimy/mir/module/faygo"
)

type Index struct {
	Id        int      `param:"<in:path> <required> <desc:ID> <range: 0:10>"`
	Title     string   `param:"<in:query> <nonzero>"`
	Paragraph []string `param:"<in:query> <name:p> <len: 1:10> <regexp: ^[\\w]*$>"`
	Cookie    string   `param:"<in:cookie> <name:faygoID>"`

	index  mir.Get `mir:"/index/:id#."`
	index2 mir.Get `mir:"/index2/:id#."`
}

func (i *Index) Serve(c *faygo.Context) error {
	if c.CookieParam("faygoID") == "" {
		c.SetCookie("faygoID", time.Now().String())
	}
	return c.JSON(200, i)
}

func main() {
	// new faygo engine
	app := faygo.New("myapp", "0.1")

	// Register handler to engine by mir
	mirE.Register(app, &Index{})

	// Start gin engine serve
	app.Run()
}

```