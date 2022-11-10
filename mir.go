// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

// Group indicator a default group for handler to register to server engine
type Group any

// Chain indicator a Handler slice used register Middleware to router by group
type Chain any

// Get indicator a GET method handler used placeholder register info in struct tag
type Get any

// Put indicator a PUT method handler used placeholder register info in struct tag
type Put any

// Post indicator a POST method handler used placeholder register info in struct tag
type Post any

// Delete indicator a DELETE method handler used placeholder register info in struct tag
type Delete any

// Head indicator a HEAD method handler used placeholder register info in struct tag
type Head any

// Patch indicator a PATCH method handler used placeholder register info in struct tag
type Patch any

// Trace indicator a TRACE method handler used placeholder register info in struct tag
type Trace any

// Connect indicator a CONNECT method handler used placeholder register info in struct tag
type Connect any

// Options indicator a OPTIONS method handler used placeholder register info in struct tag
type Options any

// Any indicator a Any method handler used placeholder register info in struct tag.
// This is mean register handler that all http.Method* include(GET, PUT, POST, DELETE,
// HEAD, PATCH, OPTIONS)
type Any any
