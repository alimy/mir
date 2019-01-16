// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package iris

import (
	"github.com/alimy/mir"
	"github.com/kataras/iris"
)

// Mir return mir.Engine interface implements instance.Used to register routes
// to iris application with struct tag string's information.
func Mir(app *iris.Application) mir.Engine {
	return &mirEngine{engine: app}
}

// Register use entries's info to register handler to iris application.
func Register(app *iris.Application, entries ...interface{}) error {
	mirE := Mir(app)
	return mir.Register(mirE, entries...)
}
