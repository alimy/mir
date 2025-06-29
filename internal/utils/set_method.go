// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"net/http"
	"sort"
)

// HttpMethodSet http method set struct
type HttpMethodSet map[string]struct{}

func (s HttpMethodSet) Add(methods ...string) {
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

func (s HttpMethodSet) List() []string {
	methods := make([]string, 0, len(s))
	for m := range s {
		methods = append(methods, m)
	}
	sort.Strings(methods)
	return methods
}
