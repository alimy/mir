// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
)

func init() {
	AddEntry(new(Site), new(Admin))
}

// Site site v1 interface info
type Site struct {
	Chain            `mir:"-"`
	Group            `mir:"v1"`
	Index            func(Get, Chain)                         `mir:"/index/"`
	AnyTopics        func(Chain)                              `mir:"/topics/"`
	Articles         func(Head, Get, Post, Chain)             `mir:"/articles/:category/"`
	Logout           func(Post)                               `mir:"/user/logout/"`
	Assets           func(Get, Context)                       `mir:"/assets"`
	AnyStaticks      func(Any, Context)                       `mir:"/staticks"`
	ManyResources    func(Get, Head, Options, Context)        `mir:"/resources"`
	MultiAttachments func(Get, Head, Options, Chain, Context) `mir:"/attachments"`
}

// Admin admin v1 interface info
type Admin struct {
	Chain   `mir:"-"`
	Group   `mir:"v1"`
	User    func(Get, Chain)             `mir:"/user/"`
	DelUser func(Delete, Chain)          `mir:"/user/"`
	Teams   func(Head, Get, Post, Chain) `mir:"/teams/:category/"`
	Quit    func(Post)                   `mir:"/user/quit/"`
}
