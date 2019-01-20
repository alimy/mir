package gin_test

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type entry struct {
	count    uint64
	Group    mir.Group `mir:"v1"`
	add      mir.Post  `mir:"/add/:id/"`
	index    mir.Get   `mir:"/index/"`
	articles mir.Get   `mir:"/articles/:category/#GetArticles"`
}

// Add handler of "/add/:id"
func (h *entry) Add(c *gin.Context) {
	body, err := h.bytesFromBody(c)
	if err != nil {
		body = []byte("")
	}
	result := strings.Join([]string{
		"Add",
		c.Params.ByName("id"),
		string(body),
	}, ":")
	c.String(http.StatusOK, result)
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *entry) Index(c *gin.Context) {
	h.count++
	c.String(http.StatusOK, "Index")
}

func (h *entry) bytesFromBody(c *gin.Context) ([]byte, error) {
	defer c.Request.Body.Close()

	buf := [256]byte{}
	result := make([]byte, 0)
	if size, err := c.Request.Body.Read(buf[:]); err == nil {
		result = append(result, buf[:size]...)
	} else if err != io.EOF {
		return nil, err
	}
	return result, nil

}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (h *entry) GetArticles(context *gin.Context) {
	context.String(http.StatusOK, "GetArticles:"+context.Params.ByName("category"))
}
