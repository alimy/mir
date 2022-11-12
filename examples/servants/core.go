package servants

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func bindAny(c *gin.Context, obj any) error {
	return c.ShouldBind(obj)
}

func renderAny(c *gin.Context, data any, err error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
