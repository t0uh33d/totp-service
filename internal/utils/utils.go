package utils

import (
	"bytes"
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func RenderTemplComponent(c *gin.Context, component templ.Component) {
	var buf bytes.Buffer
	err := component.Render(context.Background(), &buf)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering component: %v", err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
}
