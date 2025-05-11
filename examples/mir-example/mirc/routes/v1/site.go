// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v1

import (
	. "github.com/alimy/mir/v5"
)

// Site site v1 interface info
type Site struct {
	Schema           `mir:"v1,chain"`
	Index            func(Get, Chain)                         `mir:"index"`
	AnyTopics        func(Chain)                              `mir:"topics"`
	Articles         func(Head, Get, Post, Chain)             `mir:"articles/:category"`
	Logout           func(Post)                               `mir:"user/logout"`
	Assets           func(Get, Context)                       `mir:"assets/:name"`
	AnyStaticks      func(Any, Context)                       `mir:"anystaticks/:name"`
	ManyResources    func(Get, Head, Options, Context)        `mir:"resources/:name"`
	MultiAttachments func(Get, Head, Options, Chain, Context) `mir:"attachments/:name"`
}

// Admin admin v1 interface info
type Admin struct {
	Schema  `mir:"v1,chain"`
	User    func(Get, Chain)             `mir:"user"`
	DelUser func(Delete, Chain)          `mir:"user"`
	Teams   func(Head, Get, Post, Chain) `mir:"teams/:category"`
	Quit    func(Post)                   `mir:"user/quit"`
}
