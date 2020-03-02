// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v2

import "github.com/alimy/mir/v2"

// Site mir's struct tag define
type Site struct {
	Group    mir.Group `mir:"v2"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
	Category mir.Get   `mir:"/category/"`
}
