// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routers

import "github.com/alimy/mir/v2"

// Site mir's struct tag define
type Site struct {
	Chain    mir.Chain `mir:"-"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
}
