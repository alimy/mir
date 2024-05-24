package service

import (
	"fmt"

	"github.com/alimy/mir/sail/mir-example/v4/servants"

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
	s.RegisterRoute(s, servants.RegisterServants)
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
