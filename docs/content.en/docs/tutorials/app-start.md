---
weight: 25
title: "App Start"
---

## App Start:
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
 