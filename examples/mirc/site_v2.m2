// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v2

// Site v2 service
service Site(group: v2) {
	Index()     `get:"/index/"`
	Articles()  `get:"/articles/:category/"`
	Category()  `get:"/category/"`
}
