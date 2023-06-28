// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.0.0

package api

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type _binding_ interface {
	Bind(*gin.Context) mir.Error
}

type _render_ interface {
	Render(*gin.Context)
}

type _default_ interface {
	Bind(*gin.Context, any) mir.Error
	Render(*gin.Context, any, mir.Error)
}

type LogoutReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
}

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
}

type LoginReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
	Passwd    string    `json:"passwd"`
}

type LoginResp struct {
	UserInfo
	ServerInfo ServerInfo `json:"server_info"`
	JwtToken   string     `json:"jwt_token"`
}

type ServerInfo struct {
	ApiVer string `json:"api_ver"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type TweetsReq struct {
	Date string `json:"date"`
}

type TweetsResp struct {
	Tweets []Tweet `json:"tweets"`
	Total  uint32  `json:"total"`
}

type Tweet struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Site interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Logout(*LogoutReq) mir.Error
	Login(*LoginReq) (*LoginResp, mir.Error)
	PrevTweets(*TweetsReq) (*TweetsResp, mir.Error)
	NextTweets(*TweetsReq) (*TweetsResp, mir.Error)
	AnyTopics() mir.Error
	Articles() mir.Error
	Index() mir.Error

	mustEmbedUnimplementedSiteServant()
}

type SiteChain interface {
	ChainIndex() gin.HandlersChain

	mustEmbedUnimplementedSiteChain()
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site, m ...SiteChain) {
	var cc SiteChain
	if len(m) > 0 {
		cc = m[0]
	} else {
		cc = &UnimplementedSiteChain{}
	}
	router := e
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req := new(LogoutReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.Logout(req))
	})
	router.Handle("POST", "/user/login/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req := new(LoginReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.Login(req)
		s.Render(c, resp, err)
	})
	{
		h := func(c *gin.Context) {
			select {
			case <-c.Request.Context().Done():
				return
			default:
			}

			req := new(TweetsReq)
			if err := s.Bind(c, req); err != nil {
				s.Render(c, nil, err)
				return
			}
			resp, err := s.PrevTweets(req)
			s.Render(c, resp, err)
		}
		router.Handle("HEAD", "/tweets/prev/", h)
		router.Handle("GET", "/tweets/prev/", h)
		router.Handle("POST", "/tweets/prev/", h)
	}
	router.Any("/tweets/next/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req := new(TweetsReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.NextTweets(req)
		s.Render(c, resp, err)
	})
	router.Any("/topics/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.AnyTopics())
	})
	router.Handle("GET", "/articles/:category/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.Articles())
	})
	router.Handle("GET", "/index/", append(cc.ChainIndex(), func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.Index())
	})...)
}

// UnimplementedSiteServant can be embedded to have forward compatible implementations.
type UnimplementedSiteServant struct{}

func (UnimplementedSiteServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedSiteServant) Logout(req *LogoutReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Login(req *LoginReq) (*LoginResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) PrevTweets(req *TweetsReq) (*TweetsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) NextTweets(req *TweetsReq) (*TweetsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) AnyTopics() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Articles() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Index() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) mustEmbedUnimplementedSiteServant() {}

// UnimplementedSiteChain can be embedded to have forward compatible implementations.
type UnimplementedSiteChain struct{}

func (b *UnimplementedSiteChain) ChainIndex() gin.HandlersChain {
	return nil
}

func (b *UnimplementedSiteChain) mustEmbedUnimplementedSiteChain() {}
