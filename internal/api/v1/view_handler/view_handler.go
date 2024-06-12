package view_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t0uh33d/totp-service/internal/views"
)

func ViewHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "", views.Login())
}
