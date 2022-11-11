// Code generated by go-mir. DO NOT EDIT.

package v1

import (
	"errors"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
}

type ServerInfo struct {
	ApiVer string `json:"api_ver"`
}

type UserInfo struct {
	Name string `json:"name"`
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

type Site interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Index(c *gin.Context) error
	Articles(c *gin.Context) error
	Login(c *gin.Context, req *LoginReq) (*LoginResp, error)
	Logout(c *gin.Context) error

	mustEmbedUnimplementedSiteServant()
}

type SiteBinding interface {
	BindLogin(c *gin.Context) (*LoginReq, error)

	mustEmbedUnimplementedSiteBinding()
}

type SiteRender interface {
	RenderAny(c *gin.Context, data any, err error)
	RenderIndex(c *gin.Context, err error)
	RenderArticles(c *gin.Context, err error)
	RenderLogin(c *gin.Context, data *LoginResp, err error)
	RenderLogout(c *gin.Context, err error)

	mustEmbedUnimplementedSiteRender()
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site, b SiteBinding, r SiteRender) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/index/", func(c *gin.Context) {
		r.RenderIndex(c, s.Index(c))
	})
	router.Handle("GET", "/articles/:category/", func(c *gin.Context) {
		r.RenderArticles(c, s.Articles(c))
	})
	router.Handle("POST", "/user/login/", func(c *gin.Context) {
		req, err := b.BindLogin(c)
		if err != nil {
			r.RenderLogin(c, nil, err)
		}
		resp, err := s.Login(c, req)
		r.RenderLogin(c, resp, err)
	})
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		r.RenderLogout(c, s.Logout(c))
	})
}

// UnimplementedSiteServant can be embedded to have forward compatible implementations.
type UnimplementedSiteServant struct{}

// UnimplementedSiteBinding can be embedded to have forward compatible implementations.
type UnimplementedSiteBinding struct{}

// UnimplementedSiteRender can be embedded to have forward compatible implementations.
type UnimplementedSiteRender struct{}

func (UnimplementedSiteServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedSiteServant) Index(c *gin.Context) error {
	return errors.New("method Index not implemented")
}

func (UnimplementedSiteServant) Articles(c *gin.Context) error {
	return errors.New("method Index not implemented")
}

func (UnimplementedSiteServant) Login(c *gin.Context, req *LoginReq) (*LoginResp, error) {
	return nil, errors.New("method Login not implemented")
}

func (UnimplementedSiteServant) Logout(c *gin.Context) error {
	return errors.New("method Logout not implemented")
}

func (UnimplementedSiteServant) mustEmbedUnimplementedSiteServant() {}

func (UnimplementedSiteBinding) BindLogin(c *gin.Context) (*LoginReq, error) {
	return nil, errors.New("method BindLogin not implemented")
}

func (UnimplementedSiteBinding) mustEmbedUnimplementedSiteBinding() {}

func (r UnimplementedSiteRender) RenderAny(c *gin.Context, data any, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogout not implemented")
}

func (r UnimplementedSiteRender) RenderIndex(c *gin.Context, err error) {
	r.RenderAny(c, nil, err)
}

func (r UnimplementedSiteRender) RenderArticles(c *gin.Context, err error) {
	r.RenderAny(c, nil, err)
}

func (r UnimplementedSiteRender) RenderLogin(c *gin.Context, data *LoginResp, err error) {
	r.RenderAny(c, data, err)
}

func (r UnimplementedSiteRender) RenderLogout(c *gin.Context, err error) {
	r.RenderAny(c, nil, err)
}

func (UnimplementedSiteRender) mustEmbedUnimplementedSiteRender() {}
