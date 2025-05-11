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
	. "github.com/alimy/mir/v5/engine"
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