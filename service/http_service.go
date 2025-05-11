// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"net/http"
)

type BaseHttpService[T http.Handler] struct {
	baseService

	server *httpServer[T]
}

func (s *BaseHttpService[T]) RegisterRoute(srv Service, h func(e T)) {
	if h != nil {
		h(s.server.e)
	}
	s.server.addService(srv)
}

func (s *BaseHttpService[T]) OnStart() error {
	// do nothing default
	return nil
}

func (s *BaseHttpService[T]) OnStop() error {
	// do nothing default
	return nil
}

// NewBaseHttpService create a BaseHttpService instance
func NewBaseHttpService[T http.Handler](p *ServerPool[*httpServer[T]], e T, s *http.Server) *BaseHttpService[T] {
	s.Handler = e
	server := p.from(s.Addr, func() *httpServer[T] {
		return &httpServer[T]{
			baseServer: newBaseServe(),
			e:          e,
			server:     s,
		}
	})
	return &BaseHttpService[T]{
		server: server,
	}
}
