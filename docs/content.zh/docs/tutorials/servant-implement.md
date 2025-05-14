---
weight: 23
title: "接口实现"
---

## 接口实现 
```go
// file: servants/user.go

package servants

import (
	"net/http"

	"github.com/alimy/mir/v5"
	"github.com/gin-gonic/gin"
	api "github.com/alimy/mir-example/v5/mirc/auto/api/v1"
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
