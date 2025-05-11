// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"context"
	"net/http"
)

var (
	_ server = (*httpServer[*http.ServeMux])(nil)
)

// httpServer wraper for gin.engine and http.Server
type httpServer[T http.Handler] struct {
	*baseServer

	e      T
	server *http.Server
}

func (s *httpServer[T]) start() error {
	return s.server.ListenAndServe()
}

func (s *httpServer[T]) stop() error {
	return s.server.Shutdown(context.Background())
}
