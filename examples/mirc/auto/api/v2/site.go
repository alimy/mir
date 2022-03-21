// Code generated by go-mir. DO NOT EDIT.

package v2

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type Site interface {
	Index(*gin.Context)
	Articles(*gin.Context)
	Category(*gin.Context)

	mustEmbedUnimplementedSiteServant()
}

// RegisterSiteServant register Site servant to gin
func RegisterSiteServant(e *gin.Engine, s Site) {
	router := e.Group("v2")

	// register routes info to router
	router.Handle("GET", "/index/", s.Index)
	router.Handle("GET", "/articles/:category/", s.Articles)
	router.Handle("GET", "/category/", s.Category)
}

// UnimplementedSiteServant can be embedded to have forward compatible implementations.
type UnimplementedSiteServant struct {
}

func (UnimplementedSiteServant) Index(c *gin.Context) {
	c.String(http.StatusNotImplemented, "method Index not implemented")
}

func (UnimplementedSiteServant) Articles(c *gin.Context) {
	c.String(http.StatusNotImplemented, "method Articles not implemented")
}

func (UnimplementedSiteServant) Category(c *gin.Context) {
	c.String(http.StatusNotImplemented, "method Category not implemented")
}

func (UnimplementedSiteServant) mustEmbedUnimplementedSiteServant() {}
