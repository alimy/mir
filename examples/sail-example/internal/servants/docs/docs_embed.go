// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

//go:build docs
// +build docs

package docs

import (
	"github.com/alimy/mir/sail-example/v5/docs/openapi"
	"github.com/gin-gonic/gin"
)

// RouteDocs register docs asset route
func RouteDocs(e *gin.Engine) {
	e.StaticFS("/docs/openapi", openapi.NewFileSystem())
}
