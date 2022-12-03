## Mir
[![GitHub Workflow Status](https://img.shields.io/github/checks-status/alimy/mir/main?logo=github)](https://github.com/alimy/mir/actions?query=branch%3Amain)
[![GoDoc](https://godoc.org/github.com/alimy/mir?status.svg)](https://pkg.go.dev/github.com/alimy/mir/v3)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/alimy/mir)

Mir 是一套提供类似gRPC服务开发体验的快速开发RESTful API后端开发脚手架，适配多种HTTP框架，包括 [Gin](https://github.com/gin-gonic/gin), [Chi](https://github.com/go-chi/chi), [Hertz](https://github.com/cloudwego/hertz), [Echo](https://github.com/labstack/echo), [Iris](https://github.com/kataras/iris), [Fiber](https://github.com/gofiber/fiber), [Macaron](https://github.com/go-macaron/macaron), [Mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter)。  

 ![](docs/.assets/mir-arc-adapter.png) 
 
 ## 使用说明
 
 RESTful接口定义:
```go
// file: mirc/routes.go

package routes

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
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
	Chain  Chain                          `mir:"-"`
	Group  Group                          `mir:"v1"`
	Login  func(Post, LoginReq) LoginResp `mir:"/login/"`
	Logout func(Post)                     `mir:"/logout/"`
}
```

代码生成:
```go
// file: mirc/auto/api/routes.go

// Code generated by go-mir. DO NOT EDIT.
package routes

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginResp struct {
	JwtToken string `json:"jwt_token"`
}

type User interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Login(c *gin.Context, req *LoginReq) (*LoginResp, mir.Error)
	Logout(c *gin.Context) mir.Error

	mustEmbedUnimplementedUserServant()
}

type UserBinding interface {
	BindLogin(c *gin.Context) (*LoginReq, mir.Error)

	mustEmbedUnimplementedUserBinding()
}

type UserRender interface {
	RenderLogin(c *gin.Context, data *LoginResp, err mir.Error)
	RenderLogout(c *gin.Context, err mir.Error)

	mustEmbedUnimplementedUserRender()
}

// UnimplementedUserServant can be embedded to have forward compatible implementations.
type UnimplementedUserServant struct {
}

// UnimplementedSiteBinding can be embedded to have forward compatible implementations.
type UnimplementedSiteBinding struct {
	BindAny func(*gin.Context, any) mir.Error
}

// UnimplementedSiteRender can be embedded to have forward compatible implementations.
type UnimplementedSiteRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

// RegisterUserServant register User servant to gin
func RegisterUserServant(e *gin.Engine, s User, b UserBinding, r UserRender) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("POST", "/login/", func(c *gin.Context) {
		req, err := b.BindLogin(c)
		if err != nil {
			r.RenderLogin(c, nil, err)
		}
		resp, err := s.Login(c, req)
		r.RenderLogin(c, resp, err)
	})
	router.Handle("POST", "/logout/", func(c *gin.Context) {
		r.RenderLogout(c, s.Logout(c))
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

func (b UnimplementedUserBinding) BindLogin(c *gin.Context) (*LoginReq, mir.Error) {
	obj := new(LoginReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b UnimplementedUserBinding) mustEmbedUnimplementedUserBinding() {}

func (r UnimplementedUserRender) RenderLogin(c *gin.Context, data *LoginResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r UnimplementedUserRender) RenderLogout(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r UnimplementedUserRender) mustEmbedUnimplementedUserRender() {}

```

接口实现:   
```go
// file: servants/user.go

package servants

import (
	"github.com/alimy/mir-example/v3/mirc/auto/api"
)

type userSrv struct {
	api.UnimplementedUserServant
}

type userBinding struct {
	*api.UnimplementedUserBinding
}

type userRender struct {
	*api.UnimplementedUserRender
}

func newUserSrv() api.Site {
	return &userSrv{}
}

func newUserBinding() api.SiteBinding {
	return &siteBinding{
		UnimplementedSiteBinding: &api.UnimplementedSiteBinding{
			BindAny: bindAny,
		},
	}
}

func newUserRender() api.SiteRender {
	return &siteRender{
		UnimplementedSiteRender: &api.UnimplementedSiteRender{
			RenderAny: renderAny,
		},
	}
}

func bindAny(c *gin.Context, obj any) mir.Error {
	if err != c.ShouldBind(obj); err != nil {
		return mir.NewError(http.StatusBadRequest, err)
	}
	return nil
}

func renderAny(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}
```

服务注册:  
```go
// file: servants/servants.go

package servants

import (
	"github.com/alimy/mir-example/v3/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv(), newUserBinding(), newUserRender())
	
	// TODO: some other servant to register
}
```

## 使用[go-mir](https://github.com/alimy/mir)的项目
 * [examples](examples)  
[go-mir](https://github.com/alimy/mir)项目自带的demo，主要演示了如何使用[Mir](https://github.com/alimy/mir)快速进行RESTful API的后端开发.
 
