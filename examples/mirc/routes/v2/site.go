package v2

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
)

func init() {
	AddEntry(new(Site))
}

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
}

type ServerInfo struct {
	ApiVer string `json:"api_ver"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type LoginReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
	Passwd    string    `json:"passwd"`
}

type LoginResp struct {
	UserInfo
	ServerInfo ServerInfo `json:"server_info"`
	JwtToken   string     `json:"jwt_token"`
}

// Site site v2 interface info
type Site struct {
	Group    Group                          `mir:"v2"`
	Index    func(Get)                      `mir:"/index/"`
	Articles func(Get)                      `mir:"/articles/:category/"`
	Login    func(Post, LoginReq) LoginResp `mir:"/user/login/"`
	Logout   func(Post)                     `mir:"/user/logout/"`
}
