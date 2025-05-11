// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"path"
	"strings"
)

func anyPathPattern(v string) string {
	cpath := path.Clean(v)
	if !path.IsAbs(cpath) && !strings.HasPrefix(cpath, "./") && !strings.HasPrefix(cpath, "../") {
		cpath = path.Join("./", cpath)
	}
	return cpath + "./..."
}
