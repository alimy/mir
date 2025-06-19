// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/sail-example/v5/auto/api"
	v1 "github.com/alimy/mir/sail-example/v5/auto/api/v1"
	v2 "github.com/alimy/mir/sail-example/v5/auto/api/v2"
	v3 "github.com/alimy/mir/sail-example/v5/auto/api/v3"
	"github.com/gin-gonic/gin"
)

// RouteWeb register web route
func RouteWeb(e *gin.Engine) {
	api.RegisterSiteServant(e, newSiteSrv())
	v1.RegisterAdminServant(e, newAdminSrvA())
	v1.RegisterSiteServant(e, newSiteSrvA(), newSiteChainA())
	v2.RegisterSiteServant(e, newSiteSrvB())
	v3.RegisterSiteServant(e, newSiteSrvC())
}
