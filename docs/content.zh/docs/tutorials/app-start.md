---
weight: 25
title: "程序启动"
---

## 程序启动:
```go
// file: main.go

package main

import (
	"log"

	"github.com/alimy/mir-example/v4/servants"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	// register servants to gin
	servants.RegisterServants(e)

	// start servant service
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
```
 