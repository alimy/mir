// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v3.2.0

package v1

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
)

type Site interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Logout() mir.Error
	Articles() mir.Error
	AnyTopics() mir.Error
	Index() mir.Error

	mustEmbedUnimplementedSiteServant()
}

type SiteChain interface {
	ChainArticles() gin.HandlersChain
	ChainAnyTopics() gin.HandlersChain
	ChainIndex() gin.HandlersChain

	mustEmbedUnimplementedSiteChain()
}

type SiteRender interface {
	RenderLogout(*gin.Context, mir.Error)
	RenderArticles(*gin.Context, mir.Error)
	RenderAnyTopics(*gin.Context, mir.Error)
	RenderIndex(*gin.Context, mir.Error)

	mustEmbedUnimplementedSiteRender()
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site, r SiteRender, c SiteChain) {
	router := e.Group("v1")
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

		r.RenderLogout(c, s.Logout())
	})

	{
		h := append(c.ChainArticles(), func(c *gin.Context) {
			select {
			case <-c.Request.Context().Done():
				return
			default:
			}

			r.RenderArticles(c, s.Articles())
		})
		router.Handle("HEAD", "/articles/:category/", h...)
		router.Handle("POST", "/articles/:category/", h...)
		router.Handle("GET", "/articles/:category/", h...)
	}

	router.Any("/topics/", append(c.ChainAnyTopics(), func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderAnyTopics(c, s.AnyTopics())
	})...)

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

func (UnimplementedSiteServant) Chain() gin.HandlersChain {
	return nil
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

// UnimplementedSiteRender can be embedded to have forward compatible implementations.
type UnimplementedSiteRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

func (r *UnimplementedSiteRender) RenderLogout(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedSiteRender) RenderArticles(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedSiteRender) RenderAnyTopics(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedSiteRender) RenderIndex(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedSiteRender) mustEmbedUnimplementedSiteRender() {}

// UnimplementedSiteChain can be embedded to have forward compatible implementations.
type UnimplementedSiteChain struct {
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
