// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package bot

import (
	api "github.com/alimy/mir/sail/examples/v4/auto/api/bot/v1"
	"github.com/gin-gonic/gin"
)

// RouteBot register Bot route
func RouteBot(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv())
}
