package v2

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
)

func init() {
	AddEntry(new(Site))
}

type LoginReq struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginResp struct {
	JwtToken string `json:"jwt_token"`
}

// Site site v2 interface info
type Site struct {
	Group    Group                          `mir:"v2"`
	Index    func(Get)                      `mir:"/index/"`
	Articles func(Get)                      `mir:"/articles/:category/"`
	Login    func(Post, LoginReq) LoginResp `mir:"/user/login/"`
	Logout   func(Post)                     `mir:"/user/logout/"`
}
