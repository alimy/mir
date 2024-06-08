// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mir/sail-example/v4/internal/servants/bot"
	"github.com/alimy/mir/sail-example/v4/internal/servants/docs"
	"github.com/alimy/mir/sail-example/v4/internal/servants/web"
	"github.com/gin-gonic/gin"
)

// RegisterWebServants register the web servants to gin.Engine
func RegisterWebServants(e *gin.Engine) {
	web.RouteWeb(e)
}

// RegisterBotServants register the bot servants to gin.Engine
func RegisterBotServants(e *gin.Engine) {
	bot.RouteBot(e)
}

// RegisterDocsServants register the docs servants to gin.Engine
func RegisterDocsServants(e *gin.Engine) {
	docs.RouteDocs(e)
}
