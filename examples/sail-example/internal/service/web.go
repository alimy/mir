// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/alimy/mir/sail-example/v5/internal/servants"
	"github.com/alimy/mir/v5/service"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var (
	_ service.Service = (*webService)(nil)
)

type webService struct {
	*service.BaseHttpService[*gin.Engine]
	addr string
}

func (s *webService) Name() string {
	return "WebService"
}

func (s *webService) Version() string {
	return "v0.0.1"
}

func (s *webService) OnInit() error {
	s.RegisterRoute(s, servants.RegisterWebServants)
	return nil
}

func (s *webService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s", s.addr))
}

func newWebService(s *service.BaseHttpService[*gin.Engine], addr string) *webService {
	return &webService{
		BaseHttpService: s,
		addr:            addr,
	}
}
