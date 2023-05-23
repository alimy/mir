## Mir参考文档
Mir 是一套提供类似gRPC服务开发体验的快速开发RESTful API后端开发脚手架，适配多种HTTP框架，包括 [Gin](https://github.com/gin-gonic/gin), [Chi](https://github.com/go-chi/chi), [Hertz](https://github.com/cloudwego/hertz), [Echo](https://github.com/labstack/echo), [Iris](https://github.com/kataras/iris), [Fiber](https://github.com/gofiber/fiber), [Macaron](https://github.com/go-macaron/macaron), [Mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter)。 

### 使用说明
* 生成样板项目
```bash
% go get github.com/alimy/mir/mirc/v3@latest
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
% make build
```

### RESTful接口定义
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
如上示例，`file:mirc/routes.go`文件定义了`User` RESTful API接口，简要说明Mir的接口是如何定义：
* 一个独立的接口集合是定义在Golang的结构体struct中的，struct的名字将是接口集合的名字；定义后的接口需要通过`mir.AddEntry(any)`注册到Mir中以供服务接口代码的自动生成；比如这里的`User`就是一个RESTful API的集合，定义后的接口信息通过`init()`中的`AddEntry(new(User))`注册到Mir中，在Mir自动生成服务接口代码的时候将根据User中包含的信息自动生成服务接口代码；
* 接口定义的结构体struct中，每个Field都有特殊意义:  
    * Field名字表示这个服务接口自动生成代码的方法名字；
    * Field中的struct tag `mir` 定义的是接口将注册的路由url；  
        如上示例的`Login`接口将生成如下代码:  
        ```go
        // RegisterUserServant register User servant to gin
        func RegisterUserServant(e *gin.Engine, s User, b UserBinding, r UserRender) {
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

		        req, err := b.BindLogin(c)
		        if err != nil {
			        r.RenderLogin(c, nil, err)
		        }
		        resp, err := s.Login(c, req)
		        r.RenderLogin(c, resp, err)
	        })

            ...
        }
        ```
    * Field类型表示这个接口的类型，目前有三种：Chain、Group、func;   
        * `Chain` 表示这个接口集合需要生成一个获取对应HTTP引擎中间件的方法，比如gin引擎的`gin.HandlersChain`; 如上示例，使用gin引擎代码样式的代码生成器将生成如下方法:   
            ```go
            type User interface {
	        // Chain provide handlers chain for gin
	        Chain() gin.HandlersChain   
        
            ...
            }
            ```
        * `Group` 表示这是其后面的struct tag中的由`mir`标识的路径是这个接口的URL前缀/组信息；如上示例，使用gin引擎代码样式的代码生成器将生成如下注册服务方法:   
            ```go
            // RegisterUserServant register User servant to gin
            func RegisterUserServant(e *gin.Engine, s User, b UserBinding, r UserRender) {
	            router := e.Group("v1")
	            // use chain for router
	            middlewares := s.Chain()
	            router.Use(middlewares...)
                
                ...
            }
            ``` 
        * `func(...)...` 表示接口的方法定义；函数中的参数与返回值有着特殊意义:  
            * 函数可以有多个参数，也可无参数，无参数时表示将这个接口注册为所有HTTP Method handler;
            * 函数参数的类型限定为 `mir.(Get/Put/Post/Delete/Head/Patch/Trace/Connect/Options/Any/Chain)`类型、Go `struct`类型；
            * Go `struct`类型作为函数参数只能放置在最后一个参数位置，表示接口需要这个`struct`类型表示的参数类型的从http request中Binding后的结果对象作为请求参数；
            * 函数参数中的非最后一个参数，可以有多个，类型限定为`mir.(Get/Put/Post/Delete/Head/Patch/Trace/Connect/Options/Any)`类型；表示的意思是这个接口将注册为相应的HTTP Method handler，比如`mir.Post`表示将这个接口注册为 HTTP Method 为`POST` 的handler `router.Handle("POST", "/login/", func(...){...})`; `mir.Any` 表示将这个接口注册为所有HTTP Method handler;
			* 函数参数中如果有`Chain`表示这个接口有自定义的HTTP引擎中间件的方法，比如gin引擎的`gin.HandlersChain`，会与接口一起注册，目前仅支持`Gin`/`Chi`/`Echo`/`Hertz`引擎；
            * 函数返回值至多只有一个，可以没有，也可以有一个，限定类型为Go `struct`类型；
            * 函数中的最后一个参数和返回值，如果有的话，限定类型为Go `struct`类型，并且可以说与接口定义的同一个包中的结构体，也可以是其他包中的结构体； 

            如上示例，使用gin引擎样式代码生成器将生成如下代码:  
            ```go
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
            ```
### 代码生成
```go
// file: mirc/auto/api/routes.go

// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v3.1.1

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
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindLogin(c)
		if err != nil {
			r.RenderLogin(c, nil, err)
		}
		resp, err := s.Login(c, req)
		r.RenderLogin(c, resp, err)
	})
	router.Handle("POST", "/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		
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

### 接口实现 
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

### 服务注册 
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
 