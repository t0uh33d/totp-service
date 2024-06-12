package login_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/t0uh33d/totp-service/internal/utils"
	"github.com/t0uh33d/totp-service/internal/views"
)

func Login(c *gin.Context) {
	utils.RenderTemplComponent(c, views.Login())
}
