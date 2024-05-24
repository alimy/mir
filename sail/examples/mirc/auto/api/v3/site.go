// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.2.0

package v3

import (
	"context"
	"net/http"

	"github.com/alimy/mir/sail/mir-example/v4/model"
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
	RenderJsonp(*gin.Context, any, mir.Error)
}

type Site interface {
	_default_

	MultiAttachments(*gin.Context)
	ManyResources(*gin.Context)
	AnyStaticks(*gin.Context)
	Statics(*gin.Context, *model.LoginReq) mir.Error
	Assets(*gin.Context)
	Logout(context.Context) mir.Error
	Login(context.Context, *model.LoginReq) (*model.LoginResp, mir.Error)
	PrevTweets(context.Context, *model.TweetsReq) (*model.TweetsResp, mir.Error)
	NextTweets(context.Context, *model.TweetsReq) (*model.TweetsResp, mir.Error)
	Articles(context.Context) mir.Error
	Index(context.Context) mir.Error

	mustEmbedUnimplementedSiteServant()
}

type SiteChain interface {
	ChainMultiAttachments() gin.HandlersChain

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
	router := e.Group("v3")

	// register routes info to router
	{
		h := append(cc.ChainMultiAttachments(), s.MultiAttachments)
		router.Handle("GET", "/attachments/:name/", h...)
		router.Handle("HEAD", "/attachments/:name/", h...)
		router.Handle("OPTIONS", "/attachments/:name/", h...)
	}
	{
		h := s.ManyResources
		router.Handle("GET", "/resources/:name/", h)
		router.Handle("HEAD", "/resources/:name/", h)
		router.Handle("OPTIONS", "/resources/:name/", h)
	}
	router.Any("/anystaticks/:name/", s.AnyStaticks)
	router.Handle("GET", "/statics/:name/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(model.LoginReq)
		var bv _binding_ = req
		if err := bv.Bind(c); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.Statics(c, req))
	})
	router.Handle("GET", "/assets/:name/", s.Assets)
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.Logout(c.Request.Context()))
	})
	router.Handle("POST", "/user/login/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(model.LoginReq)
		var bv _binding_ = req
		if err := bv.Bind(c); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.Login(c.Request.Context(), req)
		if err != nil {
			s.Render(c, nil, err)
			return
		}
		var rv _render_ = resp
		rv.Render(c)
	})
	{
		h := func(c *gin.Context) {
			select {
			case <-c.Request.Context().Done():
				return
			default:
			}
			req := new(model.TweetsReq)
			if err := s.Bind(c, req); err != nil {
				s.Render(c, nil, err)
				return
			}
			resp, err := s.PrevTweets(c.Request.Context(), req)
			s.Render(c, resp, err)
		}
		router.Handle("GET", "/tweets/prev/", h)
		router.Handle("HEAD", "/tweets/prev/", h)
		router.Handle("POST", "/tweets/prev/", h)
	}
	router.Any("/tweets/next/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(model.TweetsReq)
		if err := s.Bind(c, req); err != nil {
			s.RenderJsonp(c, nil, err)
			return
		}
		resp, err := s.NextTweets(c.Request.Context(), req)
		s.RenderJsonp(c, resp, err)
	})
	router.Handle("GET", "/articles/:category/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.Articles(c.Request.Context()))
	})
	router.Handle("GET", "/index/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.Index(c.Request.Context()))
	})
}

// UnimplementedSiteServant can be embedded to have forward compatible implementations.
type UnimplementedSiteServant struct{}

func (UnimplementedSiteServant) MultiAttachments(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) ManyResources(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) AnyStaticks(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Statics(c *gin.Context, req *model.LoginReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Assets(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Logout(c context.Context) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Login(c context.Context, req *model.LoginReq) (*model.LoginResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) PrevTweets(c context.Context, req *model.TweetsReq) (*model.TweetsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) NextTweets(c context.Context, req *model.TweetsReq) (*model.TweetsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Articles(c context.Context) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Index(c context.Context) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) mustEmbedUnimplementedSiteServant() {}

// UnimplementedSiteChain can be embedded to have forward compatible implementations.
type UnimplementedSiteChain struct{}

func (b *UnimplementedSiteChain) ChainMultiAttachments() gin.HandlersChain {
	return nil
}

func (b *UnimplementedSiteChain) mustEmbedUnimplementedSiteChain() {}
