# Mir
[![Build Status](https://api.travis-ci.com/alimy/mir.svg?branch=master)](https://travis-ci.com/alimy/mir)
[![codecov](https://codecov.io/gh/alimy/mir/branch/master/graph/badge.svg)](https://codecov.io/gh/alimy/mir)
[![GoDoc](https://godoc.org/github.com/alimy/mir?status.svg)](https://pkg.go.dev/github.com/alimy/mir/v2)
[![sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/alimy/mir)

Mir is used for register handler to http router(eg: [Gin](https://github.com/gin-gonic/gin), [Chi](https://github.com/go-chi/chi), [Echo](https://github.com/labstack/echo), [Iris](https://github.com/kataras/iris), [Fiber](https://github.com/gofiber/fiber), [Macaron](https://github.com/go-macaron/macaron), [Mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter))
 depends on struct tag string info that defined in logic object's struct type field.
 
 ### Usage
 
 * Generate a simple template project
 
 ```
% go get github.com/alimy/mir/mirc/v2@latest
% mirc new -h
create template project

Usage:
  mirc new [flags]

Flags:
  -d, --dst string     genereted destination target directory (default ".")
  -h, --help           help for new
      --mir string     mir replace package name or place
  -p, --pkg string     project's package name (default "github.com/alimy/mir-example")
  -s, --style string   generated engine style eg: gin,chi,mux,echo,iris,fiber,macaron,httprouter (default "gin")

% mirc new -s gin -d mir-examples
% tree mir-examples
mir-examples
├── Makefile
├── README.md
├── go.mod
├── main.go
└── mirc
    ├── main.go
    └── routes
        ├── site.go
        ├── v1
        │   └── site.go
        └── v2
            └── site.go

% cd mir-examples
% make generate
 ```
 
 * Custom route info just use struct tag. eg:
 
```go
// file: mirc/routes/v1/site.go

package v1

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Site))
}

// Site mir's struct tag define
type Site struct {
	Chain    Chain `mir:"-"`
	Group    Group `mir:"v1"`
	Index    Get   `mir:"/index/"`
	Articles Get   `mir:"/articles/:category/"`
}
```

* Invoke mir's generator to generate interface. eg:

```
% cat mirc/main.go
package main

import (
	"log"

	. "github.com/alimy/mir/v2/core"
	. "github.com/alimy/mir/v2/engine"

	_ "github.com/alimy/mir/v2/examples/mirc/routes"
	_ "github.com/alimy/mir/v2/examples/mirc/routes/v1"
	_ "github.com/alimy/mir/v2/examples/mirc/routes/v2"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialMode),
		GeneratorName(GeneratorGin),
		SinkPath("./gen"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
```

* Then generate interface from routes info defined above

```go
% make generate
% cat mirc/gen/api/v1/site.go
// Code generated by go-mir. DO NOT EDIT.

package v1

import (
	"github.com/gin-gonic/gin"
)

type Site interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Index(*gin.Context)
	Articles(*gin.Context)
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/index/", s.Index)
	router.Handle("GET", "/articles/:category/", s.Articles)
}
```

* Implement api interface. eg:
```go
// file: servants/site_v1.go
package servants

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "github.com/alimy/mir/v2/examples/mirc/gen/api/v1"
)

var _ api.Site = EmptySiteV1{}

// EmptySiteV1 implement api.Site interface
type EmptySiteV1 struct{}

func (EmptySiteV1) Chain() gin.HandlersChain {
	return gin.HandlersChain{gin.Logger()}
}

func (EmptySiteV1) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data (v1)")
}

func (EmptySiteV1) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data (v1)")
}
```

* Register interface to router

```go
package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/alimy/mir/v2/examples/mirc/gen/api"
	"github.com/alimy/mir/v2/examples/mirc/gen/api/v1"
	"github.com/alimy/mir/v2/examples/mirc/gen/api/v2"
	"github.com/alimy/mir/v2/examples/servants"
)

func main() {
	e := gin.New()

	// register servants to engine
	registerServants(e)

	// start servant service
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}

func registerServants(e *gin.Engine) {
	// register default group routes
	api.RegisterSiteServant(e, servants.EmptySiteWithNoGroup{})

	// register routes for group v1
	v1.RegisterSiteServant(e, servants.EmptySiteV1{})

	// register routes for group v2
	v2.RegisterSiteServant(e, servants.EmptySiteV2{})
}
```

* Build application and run

```shell
% make run
```

### Tutorial(demo)
 * [examples](examples)  
 Just a simple exmples project for explain how to use [Mir](https://github.com/alimy/mir).
 
 * [Mirage-幻影](https://github.com/alimy/mirage)  
 Just a full web feature examples project for explain how to use [Mir](https://github.com/alimy/mir).
 
 * [mir-covid19](https://github.com/alimy/mir-covid19)  
 COVID-19 Live Updates of Tencent Health is developed to track the live updates of COVID-19, including the global pandemic trends, domestic live updates, and overseas live updates. This project is just a go version of [TH_COVID19_International](https://github.com/Tencent/TH_COVID19_International) for a guide of how to use [Mir](https://github.com/alimy/mir) to develop web application.
 
 * [zim-ms](https://github.com/alimy/zim-ms)   
Zim-ms is a demo micro-service project explain how to use [TarsGo](https://github.com/TarsCloud/TarsGo) and [go-mir](https://github.com/alimy/mir) to develop micro-service.It's easy and enjoy you heart.
