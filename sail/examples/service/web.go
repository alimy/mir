package service

import (
	"fmt"

	"github.com/alimy/mir/sail/mir-example/v4/servants"

	"github.com/alimy/mir/sail/v4/service"
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
	s.RegisterRoute(s, servants.RegisterServants)
	return nil
}

func (s *webService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s", s.addr))
}

func NewWebService(s *service.BaseHttpService[*gin.Engine], addr string) *webService {
	return &webService{
		BaseHttpService: s,
		addr:            addr,
	}
}
