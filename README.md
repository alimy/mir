## Mir
[![Go](https://github.com/alimy/mir/actions/workflows/go.yml/badge.svg)](https://github.com/alimy/mir/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/alimy/mir?status.svg)](https://pkg.go.dev/github.com/alimy/mir/v4)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/alimy/mir)

Mir 是一套提供类似gRPC服务开发体验的快速开发RESTful API后端开发脚手架，适配多种HTTP框架，包括 [Gin](https://github.com/gin-gonic/gin), [Chi](https://github.com/go-chi/chi), [Hertz](https://github.com/cloudwego/hertz), [Echo](https://github.com/labstack/echo), [Iris](https://github.com/kataras/iris), [Fiber](https://github.com/gofiber/fiber), [Macaron](https://github.com/go-macaron/macaron), [Mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter)。  

 ![](docs/static/images/mir-arc.png) 
 
 ## 使用说明
* 生成样板项目
```bash
% go install github.com/alimy/mir/mirc/v3@latest
% mirc new -h
create template project

Usage:
  mirc new [flags]

Flags:
  -d, --dst string     genereted destination target directory (default ".")
  -h, --help           help for new
      --mir string     mir replace package name or place
  -p, --pkg string     project's package name (default "github.com/alimy/mir-example")
  -s, --style string   generated engine style eg: gin,chi,mux,hertz,echo,iris,fiber,fiber-v2,macaron,httprouter (default "gin")

% mirc new -d example 
% tree example
example
.
|-- Makefile
|-- README.md
|-- go.mod
|-- go.sum
|-- main.go
|-- mirc
|   |-- auto
|   |   `-- api
|   |       |-- site.go
|   |       |-- v1
|   |       |   `-- site.go
|   |       `-- v2
|   |           `-- site.go
|   |-- gen.go
|   `-- routes
|       |-- site.go
|       |-- v1
|       |   `-- site.go
|       `-- v2
|           `-- site.go
`-- servants
    |-- core.go
    |-- servants.go
    |-- site.go
    |-- site_v1.go
    `-- site_v2.go

% cd example
% make generate
% make run
```

* RESTful接口定义:
```go
// file: mirc/routes.go

package routes

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
)

func init() {
	AddEntry(new(User))
}

type LoginReq struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginResp struct {
	JwtToken string `json:"jwt_token"`
}

// User user interface info
type User struct {
	Chain      `mir:"-"`
	Group      `mir:"v1"`
	Login  func(Post, LoginReq) LoginResp `mir:"/login/"`
	Logout func(Post)                     `mir:"/logout/"`
}
```

* 代码生成:
```go
// file: mirc/auto/api/routes.go

// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.0.0

package routes

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type _binding_ interface {
	Bind(*gin.Context) mir.Error
}

type _render_ interface {
	Render(*gin.Context)
}

type _default_ interface {
	Bind(*gin.Context, any) mir.Error
	Render(*gin.Context, any, mir.Error)
}

type LoginReq struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginResp struct {
	JwtToken string `json:"jwt_token"`
}

type User interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Login(*gin.Context, *LoginReq) (*LoginResp, mir.Error)
	Logout(*gin.Context) mir.Error

	mustEmbedUnimplementedUserServant()
}

// RegisterUserServant register User servant to gin
func RegisterUserServant(e *gin.Engine, s User) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("POST", "/login/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(LoginReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.Login(req)
		s.Render(c, resp, err)
	})
	router.Handle("POST", "/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		
		s.Render(c, nil, s.Logout(c))
	})
}

func (UnimplementedUserServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedUserServant) Login(c *gin.Context, req *LoginReq) (*LoginResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedUserServant) Logout(c *gin.Context) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedUserServant) mustEmbedUnimplementedUserServant() {}
```

* 接口实现:   
```go
// file: servants/user.go

package servants

import (
	"github.com/alimy/mir-example/v4/mirc/auto/api"
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type baseSrv struct{}

func (baseSrv) Bind(c *gin.Context, obj any) mir.Error {
	if err := c.ShouldBind(obj); err != nil {
		mir.NewError(http.StatusBadRequest, err)
	}
	return nil
}

func (baseSrv) Render(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}

type userSrv struct {
	baseSrv

	api.UnimplementedUserServant
}

func newUserSrv() api.Site {
	return &userSrv{}
}
```

* 服务注册:  
```go
// file: servants/servants.go

package servants

import (
	"github.com/alimy/mir-example/v4/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv())
	
	// TODO: some other servant to register
}
```

* 程序启动:
```go
// file: main.go

package main

import (
	"log"

	"github.com/alimy/mir-example/v4/servants"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	// register servants to gin
	servants.RegisterServants(e)

	// start servant service
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
```

### 使用[go-mir](https://github.com/alimy/mir)的项目
 * [examples](examples) - 本项目自带的demo，主要演示了如何使用[Mir](https://github.com/alimy/mir)快速进行RESTful API的后端开发。   
* [paopao-ce](https://github.com/rocboss/paopao-ce/tree/dev) - 一个清新文艺的微社区，提供类似Twiter/微博的推文分享服务。    
