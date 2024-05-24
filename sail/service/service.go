// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"log"
)

type Service interface {
	Name() string
	Version() string
	OnInit() error
	OnStart() error
	OnStop() error
}

type baseService struct{}

func (baseService) Name() string {
	return "baseService"
}

func (baseService) Version() string {
	return "v0.0.1"
}

func (baseService) String() string {
	return "baseService"
}

// MustInitService Initial service
func MustInitService(ss []Service) []Service {
	for _, s := range ss {
		if err := s.OnInit(); err != nil {
			log.Fatalf("initial %s service error: %s", s.Name(), err)
		}
	}
	return ss
}
