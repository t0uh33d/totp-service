package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	api_v1 "github.com/t0uh33d/totp-service/internal/api/v1"
	"github.com/t0uh33d/totp-service/internal/api/v1/view_handler"
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

	router.Use(api_v1.DisableCacheInDevMode)

	router.GET("/", view_handler.ViewHandler)

	router.Run(":7230")
}
