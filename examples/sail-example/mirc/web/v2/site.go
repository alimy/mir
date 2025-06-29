// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v2

import (
	. "github.com/alimy/mir/v5"
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

// Site site v2 interface info
type Site struct {
	Schema           `mir:"v2"`
	Index            func(Get, Chain)                               `mir:"index"`
	Articles         func(Get)                                      `mir:"articles/:category"`
	NextTweets       func(Any, TweetsReq) TweetsResp                `mir:"tweets/next" render:"json"`
	PrevTweets       func(Post, Get, Head, TweetsReq) TweetsResp    `mir:"tweets/prev" binding:"yaml"`
	Login            func(Post, LoginReq) LoginResp                 `mir:"user/login/" binding:"json"`
	Logout           func(Post)                                     `mir:"user/logout"`
	ImageUpload      func(Post, Context)                            `mir:"upload/image/:name"`
	FileUpload       func(Post, Chain, Context)                     `mir:"upload/file/:name"`
	SimpleUpload     func(Post, Chain, Context, LoginReq) LoginResp `mir:"upload/simple/:name/" render:"Jsonp"`
	Assets           func(Get, Context, LoginReq)                   `mir:"assets/:name"`
	AnyStaticks      func(Any, Context)                             `mir:"anystaticks/:name"`
	ManyResources    func(Get, Head, Options, Context)              `mir:"resources/:name"`
	MultiAttachments func(Get, Head, Options, Chain, Context)       `mir:"attachments/:name/" render:"XML"`
}
