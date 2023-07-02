
---
weight: 24
title: "Servant Register"
---

## Servant Register:  
```go
// file: servants/servants.go

package servants

import (
	"github.com/alimy/mir-example/v4/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

// RegisterServants register all the servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv())
	
	// TODO: some other servant to register
}
```
 