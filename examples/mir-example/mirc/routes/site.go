// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	"net/http"

	. "github.com/alimy/mir/v5"
	"github.com/gin-gonic/gin"
)

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

type LogoutReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
}

type LoginResp struct {
	UserInfo
	ServerInfo ServerInfo `json:"server_info"`
	JwtToken   string     `json:"jwt_token"`
}

type TweetsReq struct {
	Date string `json:"date"`
}

type TweetsResp struct {
	Tweets []Tweet `json:"tweets"`
	Total  uint32  `json:"total"`
}

type Tweet struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// Bind custom binding but not effect because defined in sampe package with servant interface
func (r *LoginReq) Bind(c *gin.Context) error {
	err := c.ShouldBind(r)
	if err != nil {
		return NewError(http.StatusBadRequest, err)
	}
	return nil
}

// Bind custom render but not effect because defined in sampe package with servant interface
func (r *LoginResp) Render(c *gin.Context) {
	c.String(200, "login success")
}

// Site site interface info
type Site struct {
	Schema           `mir:",chain"`
	Index            func(Get, Chain)                               `mir:"index"`
	Articles         func(Get)                                      `mir:"articles/:category"`
	AnyTopics        func()                                         `mir:"topics"`
	NextTweets       func(Any, TweetsReq) TweetsResp                `mir:"tweets/next" binding:"json"`
	PrevTweets       func(Post, Get, Head, TweetsReq) TweetsResp    `mir:"tweets/prev" render:"json"`
	Login            func(Post, LoginReq) LoginResp                 `mir:"user/login"`
	Logout           func(Post, LogoutReq)                          `mir:"user/logout"`
	ImageUpload      func(Post, Context)                            `mir:"upload/image/:name"`
	FileUpload       func(Post, Chain, Context)                     `mir:"upload/file/:name"`
	SimpleUpload     func(Post, Chain, Context, LoginReq) LoginResp `mir:"upload/simple/:name"`
	Assets           func(Get, Context, LoginReq)                   `mir:"assets/:name"`
	Statics          func(Get, Context)                             `mir:"statics/:name"`
	AnyStaticks      func(Any, Context)                             `mir:"anystaticks/:name"`
	ManyResources    func(Get, Head, Options, Context)              `mir:"resources/:name"`
	MultiAttachments func(Get, Head, Options, Chain, Context)       `mir:"attachments/:name"`
}
