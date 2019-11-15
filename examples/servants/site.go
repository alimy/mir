package servants

import (
	"net/http"

	"github.com/alimy/mir/v2/examples/mirc/gen/api/v1"
	"github.com/gin-gonic/gin"
)

var _ api.Site = EmptySite{}

// EmptySite implement api.Site interface
type EmptySite struct{}

func (EmptySite) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data")
}

func (EmptySite) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data")
}
