package servants

import (
	"github.com/alimy/mir-example/mirc/auto/api"
	"github.com/gin-gonic/gin"

	v1 "github.com/alimy/mir-example/mirc/auto/api/v1"
	v2 "github.com/alimy/mir-example/mirc/auto/api/v2"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterSiteServant(e, newSiteSrv())
	v1.RegisterSiteServant(e, newSiteV1Srv())
	v2.RegisterSiteServant(e, newSiteV2Srv())
}
