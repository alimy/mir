// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"fmt"
	"net/http"
)

var (
	_ Runtime = (*ServerPool[*httpServer[*http.ServeMux]])(nil)
)

const (
	_actOnStart byte = iota
	_actOnStop
	_actStart
	_actStop
	_actNoop
)

type server interface {
	start() error
	stop() error
	services() []Service
}

type ServerPool[T server] struct {
	servers map[string]T
}

type baseServer struct {
	ss map[string]Service
}

// NewServerPool[T] create a *ServerPool[T] instance
func NewServerPool[T server]() *ServerPool[T] {
	return &ServerPool[T]{
		servers: make(map[string]T),
	}
}

// NewHttpServerPool[T] create a *ServerPool[*httpServer[T]] instance
func NewHttpServerPool[T http.Handler]() *ServerPool[*httpServer[T]] {
	return &ServerPool[*httpServer[T]]{
		servers: make(map[string]*httpServer[T]),
	}
}

func newBaseServe() *baseServer {
	return &baseServer{
		ss: make(map[string]Service),
	}
}

func (s *baseServer) addService(srv Service) {
	if srv != nil {
		sid := srv.Name() + "@" + srv.Version()
		s.ss[sid] = srv
	}
}

func (s *baseServer) services() (ss []Service) {
	for _, s := range s.ss {
		ss = append(ss, s)
	}
	return
}

// Start start all servers
func (p *ServerPool[T]) Start(wg Launch) {
	srvSize, maxSidSize := p.checkServices()
	if srvSize < 1 {
		fmt.Fprintln(_output, "[noop] - service list is empty so do nothing")
		return
	}

	// start servers
	p.startServer(wg, maxSidSize)
}

// Stop stop all servers
func (p *ServerPool[T]) Stop() {
	srvSize, maxSidSize := p.checkServices()
	if srvSize < 1 {
		return
	}
	// stop servers
	p.stopServer(maxSidSize)
}

func (p *ServerPool[T]) from(addr string, newServer func() T) T {
	s, exist := p.servers[addr]
	if exist {
		return s
	}
	s = newServer()
	p.servers[addr] = s
	return s
}

func (p *ServerPool[T]) startServer(wg Launch, maxSidSize int) {
	for _, srv := range p.servers {
		ss := srv.services()
		if len(ss) == 0 {
			continue
		}
		startSrv := srv.start
		wg.Go(func() {
			for _, s := range ss {
				p.colorPrint(_actOnStart, s.OnStart(), maxSidSize, s)
			}
			p.colorPrint(_actStart, startSrv(), maxSidSize, ss...)
		})
	}
}

func (p *ServerPool[T]) stopServer(maxSidSize int) {
	for _, srv := range p.servers {
		ss := srv.services()
		if len(ss) < 1 {
			return
		}
		for _, s := range ss {
			p.colorPrint(_actOnStop, s.OnStop(), maxSidSize, s)
		}
		p.colorPrint(_actStop, srv.stop(), maxSidSize, ss...)
	}
}

func (p *ServerPool[T]) allServices() (ss []Service) {
	for _, srv := range p.servers {
		ss = append(ss, srv.services()...)
	}
	return
}

func (p *ServerPool[T]) checkServices() (int, int) {
	var ss []Service
	ss = append(ss, p.allServices()...)
	return len(ss), p.maxSidSize(ss)
}

func (p *ServerPool[T]) colorPrint(act byte, err error, l int, ss ...Service) {
	s := ss[0]
	switch act {
	case _actOnStart:
		if err == nil {
			fmt.Fprintf(_output, "%s [start] - %s", p.sidStr(s.Name(), s.Version(), l), s)
		} else {
			fmt.Fprintf(_output, "%s [start] - run OnStart error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
		}
	case _actOnStop:
		if err == nil {
			fmt.Fprintf(_output, "%s [stop]  - finish...\n", p.sidStr(s.Name(), s.Version(), l))
		} else {
			fmt.Fprintf(_output, "%s [stop]  - run OnStop error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
		}
	case _actStart:
		if err != nil {
			for _, s = range ss {
				fmt.Fprintf(_output, "%s [start] - starting server occurs error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
			}
		}
	case _actStop:
		if err != nil {
			for _, s = range ss {
				fmt.Fprintf(_output, "%s [stop] - stopping server occurs error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
			}
		}
	}
}

// maxSidSize max service id string length
func (p *ServerPool[T]) maxSidSize(ss []Service) int {
	length, verSize, tacitVerSize := 0, 0, len("0.0.0")
	for _, s := range ss {
		size := len(s.Name() + "@" + s.Version())
		if size > length {
			length = size
		}
		if cs := len(s.Version()); verSize < cs {
			verSize = cs
		}
	}
	return length + (verSize - tacitVerSize)
}

func (p *ServerPool[T]) sidStr(name string, version string, size int) string {
	return fmt.Sprintf(fmt.Sprintf("%%s@%%-%ds", size-len(name+version)+4), name, version)
}
