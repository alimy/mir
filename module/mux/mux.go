// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mux

import (
	"github.com/alimy/mir"
	"github.com/gorilla/mux"
)

// Mir return mir.Engine interface implements instance.
// Used to register routes to mux router with struct tag string's information.
func Mir(r *mux.Router) mir.Engine {
	return &mirEngine{engine: r}
}
