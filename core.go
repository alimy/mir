// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

// Group indicator a default group for handler to register to server engine
type Group string

// Chain indicator a Handler slice used register Middleware to router by group
type Chain interface{}

// Get indicator a GET method handler used placeholder register info in struct tag
type Get struct{}

// Put indicator a PUT method handler used placeholder register info in struct tag
type Put struct{}

// Post indicator a POST method handler used placeholder register info in struct tag
type Post struct{}

// Delete indicator a DELETE method handler used placeholder register info in struct tag
type Delete struct{}

// Head indicator a HEAD method handler used placeholder register info in struct tag
type Head struct{}

// Patch indicator a PATCH method handler used placeholder register info in struct tag
type Patch struct{}

// Trace indicator a TRACE method handler used placeholder register info in struct tag
type Trace struct{}

// Connect indicator a CONNECT method handler used placeholder register info in struct tag
type Connect struct{}

// Options indicator a OPTIONS method handler used placeholder register info in struct tag
type Options struct{}

// Any indicator a Any method handler used placeholder register info in struct tag.
// This is mean register handler that all http.Method* include(GET, PUT, POST, DELETE,
// HEAD, PATCH, OPTIONS)
type Any struct{}

// Engine register mir tag info's handler to engine.
// Other http engine router can implement this interface then use mir to register
// handler engine(eg: gin,echo,mux,httprouter)
type Engine interface {
	Register([]*TagMir) error
}
