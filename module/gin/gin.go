// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package gin

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
)

// Mir return mir.Engine interface implements instance.Used to register routes
// to gin engine with struct tag string's information.
func Mir(e *gin.Engine) mir.Engine {
	return &mirEngine{engine: e}
}

// Register use entries's info to register handler to gin engine.
func Register(e *gin.Engine, entries ...interface{}) error {
	mirE := Mir(e)
	return mir.Register(mirE, entries...)
}
