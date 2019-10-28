// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package echo

import (
	"github.com/alimy/mir"
	"github.com/labstack/echo/v4"
)

// Mir return mir.Engine interface implements instance.Used to register routes
// to echo engine with struct tag string's information.
func Mir(e *echo.Echo) mir.Engine {
	return &mirEngine{engine: e}
}

// Register use entries's info to register handler to echo engine.
func Register(e *echo.Echo, entries ...interface{}) error {
	mirE := Mir(e)
	return mir.Register(mirE, entries...)
}
