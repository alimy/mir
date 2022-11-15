// Code generated by go-mir. DO NOT EDIT.

package v1

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
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

	Index(c *gin.Context) mir.Error
	Articles(c *gin.Context) mir.Error
	Login(c *gin.Context, req *LoginReq) (*LoginResp, mir.Error)
	Logout(c *gin.Context) mir.Error

	mustEmbedUnimplementedSiteServant()
}

type SiteBinding interface {
	BindLogin(c *gin.Context) (*LoginReq, mir.Error)

	mustEmbedUnimplementedSiteBinding()
}

type SiteRender interface {
	RenderIndex(c *gin.Context, err mir.Error)
	RenderArticles(c *gin.Context, err mir.Error)
	RenderLogin(c *gin.Context, data *LoginResp, err mir.Error)
	RenderLogout(c *gin.Context, err mir.Error)

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
type UnimplementedSiteBinding struct {
	BindAny func(*gin.Context, any) mir.Error
}

// UnimplementedSiteRender can be embedded to have forward compatible implementations.
type UnimplementedSiteRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

func (UnimplementedSiteServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedSiteServant) Index(c *gin.Context) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Articles(c *gin.Context) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Login(c *gin.Context, req *LoginReq) (*LoginResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Logout(c *gin.Context) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) mustEmbedUnimplementedSiteServant() {}

func (b UnimplementedSiteBinding) BindLogin(c *gin.Context) (*LoginReq, mir.Error) {
	obj := new(LoginReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b UnimplementedSiteBinding) mustEmbedUnimplementedSiteBinding() {}

func (r UnimplementedSiteRender) RenderIndex(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r UnimplementedSiteRender) RenderArticles(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r UnimplementedSiteRender) RenderLogin(c *gin.Context, data *LoginResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r UnimplementedSiteRender) RenderLogout(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r UnimplementedSiteRender) mustEmbedUnimplementedSiteRender() {}
