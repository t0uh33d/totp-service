package login_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t0uh33d/totp-service/internal/views"
)

func Login(c *gin.Context) {
	// utils.RenderTemplComponent(c, views.Login())
	c.HTML(http.StatusOK, "", views.Login())
}
