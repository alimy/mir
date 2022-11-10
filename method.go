// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

import "net/http"

// Common HTTP methods.
const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace

	// MethodAny indicate all method above used engine register routes
	MethodAny = "ANY"
)

// HttpMethods http method list
var HttpMethods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

type MethodSet map[string]struct{}

func (s MethodSet) Add(methods ...string) {
	for _, method := range methods {
		switch method {
		case http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace:
			s[method] = struct{}{}
		}
	}
}

func (s MethodSet) List() []string {
	methods := make([]string, 0, len(s))
	for m := range s {
		methods = append(methods, m)
	}
	return methods
}
