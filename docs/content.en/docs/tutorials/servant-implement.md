---
weight: 23
title: "Servant Implement"
---

## Servant Implement:   
```go
// file: servants/user.go

package servants

import (
	"github.com/alimy/mir-example/v4/mirc/auto/api"
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type baseSrv struct{}

func (baseSrv) Bind(c *gin.Context, obj any) mir.Error {
	if err := c.ShouldBind(obj); err != nil {
		mir.NewError(http.StatusBadRequest, err)
	}
	return nil
}

func (baseSrv) Render(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}

type userSrv struct {
	baseSrv

	api.UnimplementedUserServant
}

func newUserSrv() api.Site {
	return &userSrv{}
}
```