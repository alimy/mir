// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v2

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
)

func init() {
	Entry[Site]()
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
	Group      `mir:"v2"`
	Index      func(Get, Chain)                            `mir:"/index/"`
	Articles   func(Get)                                   `mir:"/articles/:category/"`
	NextTweets func(Any, TweetsReq) TweetsResp             `mir:"/tweets/next/"`
	PrevTweets func(Post, Get, Head, TweetsReq) TweetsResp `mir:"/tweets/prev/"`
	Login      func(Post, LoginReq) LoginResp              `mir:"/user/login/"`
	Logout     func(Post)                                  `mir:"/user/logout/"`
}
