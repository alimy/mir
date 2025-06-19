// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"path"
	"strings"
)

func anyPathPatterns(schemaPath []string) (res []string) {
	var cpath string
	for _, it := range schemaPath {
		cpath = path.Clean(it)
		if !path.IsAbs(cpath) && !strings.HasPrefix(cpath, "./") && !strings.HasPrefix(cpath, "../") {
			cpath = "./" + cpath
		}
		res = append(res, cpath+"/...")
	}
	return
}
