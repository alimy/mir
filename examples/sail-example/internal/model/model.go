// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package model

import (
	"net/http"

	"github.com/alimy/mir/v5"
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
	AgentInfo  AgentInfo `json:"agent_info"`
	Name       string    `json:"name"`
	Passwd     string    `json:"passwd"`
	TempUserID string    `json:"-"`
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

func (r *TweetsReq) Precheck(c *gin.Context) error {
	if origin := c.GetHeader("Origin"); origin != "example.com" {
		return mir.Errorln(http.StatusForbidden, http.StatusText(http.StatusForbidden))
	}
	return nil
}

func (r *TweetsReq) Verify() error {
	if len(r.Date) == 0 {
		return mir.Errorln(http.StatusBadRequest, "empty date is not allowed")
	}
	return nil
}

func (r *LoginReq) Bind(c *gin.Context) error {
	err := c.ShouldBind(r)
	if err != nil {
		return mir.NewError(http.StatusBadRequest, err)
	}
	return nil
}

func (r *LoginReq) Adjust() {
	r.TempUserID = r.AgentInfo.UserAgent + r.Name
}

func (r *LoginResp) Render(c *gin.Context) {
	c.String(200, "login success")
}
