// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.0.0

package v1

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

type Admin interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	quit() mir.Error
	Teams() mir.Error
	DelUser() mir.Error
	User() mir.Error

	mustEmbedUnimplementedAdminServant()
}

type AdminChain interface {
	ChainTeams() gin.HandlersChain
	ChainDelUser() gin.HandlersChain
	ChainUser() gin.HandlersChain

	mustEmbedUnimplementedAdminChain()
}

// RegisterAdminServant register Admin servant to gin
func RegisterAdminServant(e *gin.Engine, s Admin, m ...AdminChain) {
	var cc AdminChain
	if len(m) > 0 {
		cc = m[0]
	} else {
		cc = &UnimplementedAdminChain{}
	}
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("POST", "/user/quit/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.quit())
	})
	{
		h := append(cc.ChainTeams(), func(c *gin.Context) {
			select {
			case <-c.Request.Context().Done():
				return
			default:
			}

			s.Render(c, nil, s.Teams())
		})
		router.Handle("GET", "/teams/:category/", h...)
		router.Handle("HEAD", "/teams/:category/", h...)
		router.Handle("POST", "/teams/:category/", h...)
	}
	router.Handle("DELETE", "/user/", append(cc.ChainDelUser(), func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.DelUser())
	})...)
	router.Handle("GET", "/user/", append(cc.ChainUser(), func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.User())
	})...)
}

// UnimplementedAdminServant can be embedded to have forward compatible implementations.
type UnimplementedAdminServant struct{}

func (UnimplementedAdminServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedAdminServant) quit() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAdminServant) Teams() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAdminServant) DelUser() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAdminServant) User() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAdminServant) mustEmbedUnimplementedAdminServant() {}

// UnimplementedAdminChain can be embedded to have forward compatible implementations.
type UnimplementedAdminChain struct{}

func (b *UnimplementedAdminChain) ChainTeams() gin.HandlersChain {
	return nil
}

func (b *UnimplementedAdminChain) ChainDelUser() gin.HandlersChain {
	return nil
}

func (b *UnimplementedAdminChain) ChainUser() gin.HandlersChain {
	return nil
}

func (b *UnimplementedAdminChain) mustEmbedUnimplementedAdminChain() {}
