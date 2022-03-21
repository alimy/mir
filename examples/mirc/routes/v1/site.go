package v1

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Site))
}

// Site site v1 interface info
type Site struct {
	Chain    Chain `mir:"-"`
	Group    Group `mir:"v1"`
	Index    Get   `mir:"/index/"`
	Articles Get   `mir:"/articles/:category/"`
}
