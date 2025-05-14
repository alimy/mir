---
weight: 21
title: "Interface Define"
---

## Interface Define
```go
// file: mirc/routes.go

package routes

import (
	. "github.com/alimy/mir/v5"
)

type LoginReq struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginResp struct {
	JwtToken string `json:"jwt_token"`
}

// User user interface info
type User struct {
	Schema                                `mir:"v1,chain"`
	Login  func(Post, LoginReq) LoginResp `mir:"/login/"`
	Logout func(Post)                     `mir:"/logout/"`
}
```