// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

//mir:syntax v0.1-alpha.1

package v1

// Site v1 service
service Site(group: v1, chain: _) {
	Index()     `get:"/index/"`
	Articles()  `get:"/articles/:category/"`
}
