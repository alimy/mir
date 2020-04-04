// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v2

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Site))
}

// Site mir's struct tag define
type Site struct {
	Group    Group `mir:"v2"`
	Index    Get   `mir:"/index/"`
	Articles Get   `mir:"/articles/:category/"`
	Category Get   `mir:"/category/"`
}
