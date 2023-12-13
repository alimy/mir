// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.1.0

package v1

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type Site interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	MultiAttachments(*gin.Context)
	ManyResources(*gin.Context)
	AnyStaticks(*gin.Context)
	Assets(*gin.Context)
	Logout() mir.Error
	Articles() mir.Error
	AnyTopics() mir.Error
	Index() mir.Error

	mustEmbedUnimplementedSiteServant()
}

type SiteChain interface {
	ChainMultiAttachments() gin.HandlersChain
	ChainArticles() gin.HandlersChain
	ChainAnyTopics() gin.HandlersChain
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
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	{
		h := append(cc.ChainMultiAttachments(), s.MultiAttachments)
		router.Handle("OPTIONS", "/attachments", h...)
		router.Handle("HEAD", "/attachments", h...)
		router.Handle("GET", "/attachments", h...)
	}
	{
		h := s.ManyResources
		router.Handle("OPTIONS", "/resources", h)
		router.Handle("HEAD", "/resources", h)
		router.Handle("GET", "/resources", h)
	}
	router.Any("/staticks", s.AnyStaticks)
	router.Handle("GET", "/assets", s.Assets)
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.Logout())
	})
	{
		h := append(cc.ChainArticles(), func(c *gin.Context) {
			select {
			case <-c.Request.Context().Done():
				return
			default:
			}

			s.Render(c, nil, s.Articles())
		})
		router.Handle("POST", "/articles/:category/", h...)
		router.Handle("GET", "/articles/:category/", h...)
		router.Handle("HEAD", "/articles/:category/", h...)
	}
	router.Any("/topics/", append(cc.ChainAnyTopics(), func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.AnyTopics())
	})...)
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

func (UnimplementedSiteServant) MultiAttachments(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) ManyResources(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) AnyStaticks(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Assets(c *gin.Context) {
	c.String(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Logout() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Articles() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) AnyTopics() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) Index() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedSiteServant) mustEmbedUnimplementedSiteServant() {}

// UnimplementedSiteChain can be embedded to have forward compatible implementations.
type UnimplementedSiteChain struct{}

func (b *UnimplementedSiteChain) ChainMultiAttachments() gin.HandlersChain {
	return nil
}

func (b *UnimplementedSiteChain) ChainArticles() gin.HandlersChain {
	return nil
}

func (b *UnimplementedSiteChain) ChainAnyTopics() gin.HandlersChain {
	return nil
}

func (b *UnimplementedSiteChain) ChainIndex() gin.HandlersChain {
	return nil
}

func (b *UnimplementedSiteChain) mustEmbedUnimplementedSiteChain() {}
