// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/alimy/mir/sail/examples/v4/internal/servants"
	"github.com/alimy/mir/sail/v4/service"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var (
	_ service.Service = (*botService)(nil)
)

type botService struct {
	*service.BaseHttpService[*gin.Engine]
	addr string
}

func (s *botService) Name() string {
	return "BotService"
}

func (s *botService) Version() string {
	return "v0.0.1"
}

func (s *botService) OnInit() error {
	s.RegisterRoute(s, servants.RegisterBotServants)
	return nil
}

func (s *botService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s", s.addr))
}

func newBotService(s *service.BaseHttpService[*gin.Engine], addr string) *botService {
	return &botService{
		BaseHttpService: s,
		addr:            addr,
	}
}
