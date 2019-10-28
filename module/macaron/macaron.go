// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package macaron

import (
	"github.com/alimy/mir"
	"gopkg.in/macaron.v1"
)

// Mir return mir.Engine interface implements instance.Used to register routes
// to Macaron engine with struct tag string's information.
func Mir(m *macaron.Macaron) mir.Engine {
	return &mirEngine{engine: m}
}

// Register use entries's info to register handler to Macaron engine.
func Register(m *macaron.Macaron, entries ...interface{}) error {
	mirE := Mir(m)
	return mir.Register(mirE, entries...)
}
