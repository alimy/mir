// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

type (
	// Schema is the restful Interface.
	// It can be embedded in end-user schemas as follows:
	//
	//	type WebAPI struct {
	//		mir.Schema
	//	}
	//
	Schema interface {
		mustEmbedSchema()
	}

	// Group indicator a default group for handler to register to server engine
	Group any

	// Chain indicator a Handler slice used register Middleware to router by group
	Chain any

	// Context indicator a handler that just use engine's context like use gin engine
	// the handler will just take a function func(c *gin.Context) as handler
	Context any

	// Get indicator a GET method handler used placeholder register info in struct tag
	Get any

	// Put indicator a PUT method handler used placeholder register info in struct tag
	Put any

	// Post indicator a POST method handler used placeholder register info in struct tag
	Post any

	// Delete indicator a DELETE method handler used placeholder register info in struct tag
	Delete any

	// Head indicator a HEAD method handler used placeholder register info in struct tag
	Head any

	// Patch indicator a PATCH method handler used placeholder register info in struct tag
	Patch any

	// Trace indicator a TRACE method handler used placeholder register info in struct tag
	Trace any

	// Connect indicator a CONNECT method handler used placeholder register info in struct tag
	Connect any

	// Options indicator a OPTIONS method handler used placeholder register info in struct tag
	Options any

	// Any indicator a Any method handler used placeholder register info in struct tag.
	// This is mean register handler that all http.Method* include(GET, PUT, POST, DELETE,
	// HEAD, PATCH, OPTIONS)
	Any any
)
