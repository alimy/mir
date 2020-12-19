// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package api

// Site service
service Site(_ Chain) {
	Index    Get   `/index/`
	Articles Get   `/articles/:category/`
}
