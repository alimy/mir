// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mir/v2/examples/mirc/auto/api"
	"github.com/gin-gonic/gin"

	v1 "github.com/alimy/mir/v2/examples/mirc/auto/api/v1"
	v2 "github.com/alimy/mir/v2/examples/mirc/auto/api/v2"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterSiteServant(e, newSiteSrv())
	v1.RegisterSiteServant(e, newSiteV1Srv())
	v2.RegisterSiteServant(e, newSiteV2Srv())
}
