// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v3.2.0

package v2

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
	Passwd    string    `json:"passwd"`
}

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
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
	Logout() mir.Error
	Login(*LoginReq) (*LoginResp, mir.Error)
	PrevTweets(*TweetsReq) (*TweetsResp, mir.Error)
	NextTweets(*TweetsReq) (*TweetsResp, mir.Error)
	Articles() mir.Error
	Index() mir.Error

	mustEmbedUnimplementedSiteServant()
}

type SiteBinding interface {
	BindLogin(*gin.Context) (*LoginReq, mir.Error)
	BindPrevTweets(*gin.Context) (*TweetsReq, mir.Error)
	BindNextTweets(*gin.Context) (*TweetsReq, mir.Error)

	mustEmbedUnimplementedSiteBinding()
}

type SiteChain interface {
	ChainIndex() gin.HandlersChain

	mustEmbedUnimplementedSiteChain()
}

type SiteRender interface {
	RenderLogout(*gin.Context, mir.Error)
	RenderLogin(*gin.Context, *LoginResp, mir.Error)
	RenderPrevTweets(*gin.Context, *TweetsResp, mir.Error)
	RenderNextTweets(*gin.Context, *TweetsResp, mir.Error)
	RenderArticles(*gin.Context, mir.Error)
	RenderIndex(*gin.Context, mir.Error)

	mustEmbedUnimplementedSiteRender()
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site, b SiteBinding, r SiteRender, c SiteChain) {
	router := e.Group("v2")

	// register routes info to router
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderLogout(c, s.Logout())
	})

	router.Handle("POST", "/user/login/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindLogin(c)
		if err != nil {
			r.RenderLogin(c, nil, err)
			return
		}
		resp, err := s.Login(req)
		r.RenderLogin(c, resp, err)
	})

	{
		h := func(c *gin.Context) {
			select {
			case <-c.Request.Context().Done():
				return
			default:
			}

			req, err := b.BindPrevTweets(c)
			if err != nil {
				r.RenderPrevTweets(c, nil, err)
				return
			}
			resp, err := s.PrevTweets(req)
			r.RenderPrevTweets(c, resp, err)
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

		req, err := b.BindNextTweets(c)
		if err != nil {
			r.RenderNextTweets(c, nil, err)
			return
		}
		resp, err := s.NextTweets(req)
		r.RenderNextTweets(c, resp, err)
	})

	router.Handle("GET", "/articles/:category/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderArticles(c, s.Articles())
	})

	router.Handle("GET", "/index/", append(c.ChainIndex(), func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderIndex(c, s.Index())
	})...)

}

// UnimplementedSiteServant can be embedded to have forward compatible implementations.
type UnimplementedSiteServant struct {
}

func (UnimplementedSiteServant) Logout() mir.Error {
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

func (UnimplementedSiteServant) Articles() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Index() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) mustEmbedUnimplementedSiteServant() {}

// UnimplementedSiteRender can be embedded to have forward compatible implementations.
type UnimplementedSiteRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

func (r *UnimplementedSiteRender) RenderLogout(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedSiteRender) RenderLogin(c *gin.Context, data *LoginResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedSiteRender) RenderPrevTweets(c *gin.Context, data *TweetsResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedSiteRender) RenderNextTweets(c *gin.Context, data *TweetsResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedSiteRender) RenderArticles(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedSiteRender) RenderIndex(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedSiteRender) mustEmbedUnimplementedSiteRender() {}

// UnimplementedSiteBinding can be embedded to have forward compatible implementations.
type UnimplementedSiteBinding struct {
	BindAny func(*gin.Context, any) mir.Error
}

func (b *UnimplementedSiteBinding) BindLogin(c *gin.Context) (*LoginReq, mir.Error) {
	obj := new(LoginReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedSiteBinding) BindPrevTweets(c *gin.Context) (*TweetsReq, mir.Error) {
	obj := new(TweetsReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedSiteBinding) BindNextTweets(c *gin.Context) (*TweetsReq, mir.Error) {
	obj := new(TweetsReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedSiteBinding) mustEmbedUnimplementedSiteBinding() {}

// UnimplementedSiteChain can be embedded to have forward compatible implementations.
type UnimplementedSiteChain struct {
}

func (b *UnimplementedSiteChain) ChainIndex() gin.HandlersChain {
	return nil
}

func (b *UnimplementedSiteChain) mustEmbedUnimplementedSiteChain() {}
