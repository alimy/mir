// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"testing"

	. "github.com/alimy/mir/v3"
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

type site struct {
	Chain    Chain                          `mir:"-"`
	Index    func(Get)                      `mir:"/index/"`
	Articles func(Get)                      `mir:"/articles/:category/"`
	Login    func(Post, LoginReq) LoginResp `mir:"/user/login/"`
	Logout   func(Post)                     `mir:"/user/logout/"`
}

type siteV1 struct {
	Chain    Chain                          `mir:"-"`
	Group    Group                          `mir:"v1"`
	Index    func(Get)                      `mir:"/index/"`
	Articles func(Get)                      `mir:"/articles/:category/"`
	Login    func(Post, LoginReq) LoginResp `mir:"/user/login/"`
	Logout   func(Post)                     `mir:"/user/logout/"`
}

type siteV2 struct {
	Group    Group                          `mir:"v2"`
	Index    func(Get)                      `mir:"/index/"`
	Articles func(Get)                      `mir:"/articles/:category/"`
	Login    func(Post, LoginReq) LoginResp `mir:"/user/login/"`
	Logout   func(Post)                     `mir:"/user/logout/"`
}

func TestMirParser_Parse(t *testing.T) {
	p := &mirParser{tagName: defaultTag}

	entries := []any{
		new(site),
		new(siteV1),
		new(siteV2),
	}
	ds, err := p.Parse(entries)
	if err != nil {
		t.Errorf("want nil error but got: %s", err)
	}
	if len(ds) != 3 {
		t.Errorf("want 3 item but got: %d", len(ds))
	}

	iface, exist := ds.Get("")
	if !exist || len(iface) != 1 {
		t.Error("want a correct iface but not")
	}
	site := iface["site"]
	if site == nil || len(site.Fields) != 4 {
		t.Error("want a correct iface but not")
	}

	iface, exist = ds.Get("v1")
	if !exist || len(iface) != 1 {
		t.Error("want a correct iface but not")
	}
	site = iface["siteV1"]
	if site == nil || len(site.Fields) != 4 {
		t.Error("want a correct iface but not")
	}

	iface, exist = ds.Get("v2")
	if !exist || len(iface) != 1 {
		t.Error("want a correct iface but not")
	}
	site = iface["siteV2"]
	if site == nil || len(site.Fields) != 4 {
		t.Error("want a correct iface but not")
	}
}
