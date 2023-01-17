// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mir-example/v3/mirc/auto/api"
	v1 "github.com/alimy/mir-example/v3/mirc/auto/api/v1"
	v2 "github.com/alimy/mir-example/v3/mirc/auto/api/v2"
	"github.com/gin-gonic/gin"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterSiteServant(e, newSiteSrv(), newSiteBinding(), newSiteRender())
	v1.RegisterSiteServant(e, newSiteV1Srv(), newSiteV1Render())
	v2.RegisterSiteServant(e, newSiteV2Srv(), newSiteV2Binding(), newSiteV2Render())
}
