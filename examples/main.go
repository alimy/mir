package main

import (
	"github.com/alimy/mir/v2/examples/mirc/gen/mir"
	"github.com/alimy/mir/v2/examples/servants"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()
	// register site servant
	mir.RegisterSiteServant(e, servants.EmptySite{})
	// start servant service
	mir.Serve()
}
