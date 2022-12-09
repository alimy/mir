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
