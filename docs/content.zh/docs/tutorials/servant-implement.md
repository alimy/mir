---
weight: 23
title: "接口实现"
---

## 接口实现 
```go
// file: servants/user.go

package servants

import (
	"github.com/alimy/mir-example/v4/mirc/auto/api"
)

type userSrv struct {
	api.UnimplementedUserServant
}

type userBinding struct {
	*api.UnimplementedUserBinding
}

type userRender struct {
	*api.UnimplementedUserRender
}

func newUserSrv() api.Site {
	return &userSrv{}
}

func newUserBinding() api.SiteBinding {
	return &siteBinding{
		UnimplementedSiteBinding: &api.UnimplementedSiteBinding{
			BindAny: bindAny,
		},
	}
}

func newUserRender() api.SiteRender {
	return &siteRender{
		UnimplementedSiteRender: &api.UnimplementedSiteRender{
			RenderAny: renderAny,
		},
	}
}

func bindAny(c *gin.Context, obj any) mir.Error {
	if err != c.ShouldBind(obj); err != nil {
		return mir.NewError(http.StatusBadRequest, err)
	}
	return nil
}

func renderAny(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}
```
