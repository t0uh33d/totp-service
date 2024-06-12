package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	v1 "github.com/t0uh33d/totp-service/internal/api/v1"
	"github.com/t0uh33d/totp-service/internal/api/v1/login_handler"
	"github.com/t0uh33d/totp-service/internal/config"
	"github.com/t0uh33d/totp-service/internal/tmpl_renderer"
)

func main() {
	// loading the config
	config.LoadConfig()

	// db.ConnectToDatabase()

	// defer db.SQLx.Close()

	router := gin.Default()

	htmlRenderer := router.HTMLRender
	router.HTMLRender = &tmpl_renderer.HTMLTemplRenderer{FallbackHtmlRenderer: htmlRenderer}

	v1 := router.Group("/v1", v1.DisableCacheInDevMode)

	v1.GET("/login", login_handler.Login)

	router.Run(":7230")
}
