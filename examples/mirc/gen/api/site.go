// Code generated by go-mir. DO NOT EDIT.

package api

import (
	"github.com/gin-gonic/gin"
)

type Site interface {
	Index(c *gin.Context)
	Articles(c *gin.Context)
}