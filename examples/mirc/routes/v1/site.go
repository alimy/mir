// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
)

func init() {
	AddEntry(new(Site))
}

// Site site v1 interface info
type Site struct {
	Chain    Chain                 `mir:"-"`
	Group    Group                 `mir:"v1"`
	Index    func(Get)             `mir:"/index/"`
	Articles func(Head, Get, Post) `mir:"/articles/:category/"`
}
