---
weight: 24
title: "服务注册"
---

## 服务注册 
```go
// file: servants/servants.go

package servants

import (
	api "github.com/alimy/mir-example/v4/mirc/auto/api/v1"
	"github.com/gin-gonic/gin"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv())
	
	// TODO: some other servant to register
}
```
 