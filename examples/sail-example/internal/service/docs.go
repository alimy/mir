// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
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
	_ service.Service = (*docsService)(nil)
)

type docsService struct {
	*service.BaseHttpService[*gin.Engine]
	addr string
}

func (s *docsService) Name() string {
	return "DocsService"
}

func (s *docsService) Version() string {
	return "v0.0.1"
}

func (s *docsService) OnInit() error {
	s.RegisterRoute(s, servants.RegisterBotServants)
	return nil
}

func (s *docsService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s", s.addr))
}

func newDocsService(s *service.BaseHttpService[*gin.Engine], addr string) *docsService {
	return &docsService{
		BaseHttpService: s,
		addr:            addr,
	}
}
