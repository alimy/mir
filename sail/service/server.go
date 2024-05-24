// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sourcegraph/conc"
)

const (
	_actOnStart byte = iota
	_actOnStop
	_actStart
	_actStop
)

type server interface {
	start() error
	stop() error
	services() []Service
}

type serverPool[T server] struct {
	servers map[string]T
}

type baseServer struct {
	ss map[string]Service
}

// NewServerPool[T] create a ServerPool[T] instance
func NewServerPool[T server]() *serverPool[T] {
	return &serverPool[T]{
		servers: make(map[string]T),
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
func (p *serverPool[T]) Start(wg *conc.WaitGroup) {
	srvSize, maxSidSize := p.checkServices()
	if srvSize < 1 {
		return
	}

	// start servers
	p.startServer(wg, maxSidSize)
}

// Stop stop all servers
func (p *serverPool[T]) Stop() {
	srvSize, maxSidSize := p.checkServices()
	if srvSize < 1 {
		return
	}
	// stop servers
	p.stopServer(maxSidSize)
}

func (p *serverPool[T]) from(addr string, newServer func() T) T {
	s, exist := p.servers[addr]
	if exist {
		return s
	}
	s = newServer()
	p.servers[addr] = s
	return s
}

func (p *serverPool[T]) startServer(wg *conc.WaitGroup, maxSidSize int) {
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

func (p *serverPool[T]) stopServer(maxSidSize int) {
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

func (p *serverPool[T]) allServices() (ss []Service) {
	for _, srv := range p.servers {
		ss = append(ss, srv.services()...)
	}
	return
}

func (p *serverPool[T]) checkServices() (int, int) {
	var ss []Service
	ss = append(ss, p.allServices()...)
	return len(ss), p.maxSidSize(ss)
}

func (p *serverPool[T]) colorPrint(act byte, err error, l int, ss ...Service) {
	s := ss[0]
	switch act {
	case _actOnStart:
		if err == nil {
			fmt.Fprintf(color.Output, "%s [start] - %s", p.sidStr(s.Name(), s.Version(), l), s)
		} else {
			fmt.Fprintf(color.Output, "%s [start] - run OnStart error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
		}
	case _actOnStop:
		if err == nil {
			fmt.Fprintf(color.Output, "%s [stop]  - finish...\n", p.sidStr(s.Name(), s.Version(), l))
		} else {
			fmt.Fprintf(color.Output, "%s [stop]  - run OnStop error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
		}
	case _actStart:
		if err != nil {
			for _, s = range ss {
				fmt.Fprintf(color.Output, "%s [start] - starting server occurs error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
			}
		}
	case _actStop:
		if err != nil {
			for _, s = range ss {
				fmt.Fprintf(color.Output, "%s [stop] - stopping server occurs error: %s\n", p.sidStr(s.Name(), s.Version(), l), err)
			}
		}
	}
}

// maxSidSize max service id string length
func (p *serverPool[T]) maxSidSize(ss []Service) int {
	length := 0
	for _, s := range ss {
		size := len(s.Name() + "@" + s.Version())
		if size > length {
			length = size
		}
	}
	return length
}

func (p *serverPool[T]) sidStr(name string, version string, size int) string {
	return fmt.Sprintf(fmt.Sprintf("%%s@%%-%ds", size-len(name+version)+4), name, version)
}
