// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v1

import (
	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/engine"
)

func init() {
	engine.AddEntry(new(Site))
}

// Site mir's struct tag define
type Site struct {
	Chain    mir.Chain `mir:"-"`
	Group    mir.Group `mir:"v1"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
}
