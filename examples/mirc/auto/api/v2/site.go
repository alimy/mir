// Code generated by go-mir. DO NOT EDIT.

package v2

import (
	"errors"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type LoginReq struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginResp struct {
	JwtToken string `json:"jwt_token"`
}

type Site interface {
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
	RenderIndex(c *gin.Context, err error)
	RenderArticles(c *gin.Context, err error)
	RenderLogin(c *gin.Context, data *LoginResp, err error)
	RenderLogout(c *gin.Context, err error)

	mustEmbedUnimplementedSiteRender()
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site, b SiteBinding, r SiteRender) {
	router := e.Group("v2")

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

func (UnimplementedSiteRender) RenderIndex(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogout not implemented")
}

func (UnimplementedSiteRender) RenderArticles(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogout not implemented")
}

func (UnimplementedSiteRender) RenderLogin(c *gin.Context, data *LoginResp, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogin not implemented")
}

func (UnimplementedSiteRender) RenderLogout(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogout not implemented")
}

func (UnimplementedSiteRender) mustEmbedUnimplementedSiteRender() {}
