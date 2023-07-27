---
weight: 21
title: "接口定义"
---

## 接口定义
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
			* 函数参数中如果有`Chain`表示这个接口有自定义的HTTP引擎中间件的方法，比如gin引擎的`gin.HandlersChain`，会与接口一起注册，目前仅支持`Gin`/`Chi`/`Echo`/`Hertz`/`Iris`引擎；
            * 函数返回值至多只有一个，可以没有，也可以有一个，限定类型为Go `struct`类型；
            * 函数中的最后一个参数和返回值，如果有的话，限定类型为Go `struct`类型，并且可以说与接口定义的同一个包中的结构体，也可以是其他包中的结构体； 

            如上示例，使用gin引擎样式代码生成器将生成如下代码:  
            ```go
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

	            Login(c *gin.Context, req *LoginReq) (*LoginResp, mir.Error)
	            Logout(c *gin.Context) mir.Error

	            mustEmbedUnimplementedUserServant()
            }
            ```