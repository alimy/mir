// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
)

func init() {
	AddEntry(new(Admin))
}

// Admin admin v1 interface info
type Admin struct {
	Chain   `mir:"-"`
	Group   `mir:"v1"`
	User    func(Get, Chain)             `mir:"/user/"`
	DelUser func(Delete, Chain)          `mir:"/user/"`
	Teams   func(Head, Get, Post, Chain) `mir:"/teams/:category/"`
	quit    func(Post)                   `mir:"/user/quit/"`
}
