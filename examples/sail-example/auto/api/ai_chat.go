// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir 5.2

package api

import (
	"context"
	"net/http"

	"github.com/alimy/mir/v5"
	"github.com/gin-gonic/gin"
)

type _binding_ interface {
	Bind(*gin.Context) error
}

type _render_ interface {
	Render(*gin.Context)
}

type _default_ interface {
	Bind(*gin.Context, any) error
	BindJson(*gin.Context, any) error
	Render(*gin.Context, any, error)
	RenderJson(*gin.Context, any, error)
}

type LlmsResp struct {
	Content string
}

type VersionResp struct {
	Ver string
}

type AiChat interface {
	_default_

	Llms(context.Context) (*LlmsResp, error)
	Version(context.Context) (*VersionResp, error)

	mustEmbedUnimplementedAiChatServant()
}

// RegisterAiChatServant register AiChat servant to gin
func RegisterAiChatServant(e *gin.Engine, s AiChat) {
	router := e

	// register routes info to router
	router.Handle("GET", "llms.txt", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		resp, err := s.Llms(c.Request.Context())
		s.Render(c, resp, err)
	})
	router.Handle("GET", "version", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		resp, err := s.Version(c.Request.Context())
		s.Render(c, resp, err)
	})
}

// UnimplementedAiChatServant can be embedded to have forward compatible implementations.
type UnimplementedAiChatServant struct{}

func (UnimplementedAiChatServant) Llms(c context.Context) (*LlmsResp, error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAiChatServant) Version(c context.Context) (*VersionResp, error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAiChatServant) mustEmbedUnimplementedAiChatServant() {}
