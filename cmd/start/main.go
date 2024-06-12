package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/t0uh33d/totp-service/internal/api/v1/login_handler"
	"github.com/t0uh33d/totp-service/internal/config"
)

func main() {
	// loading the config
	config.LoadConfig()

	// db.ConnectToDatabase()

	// defer db.SQLx.Close()

	router := gin.Default()

	v1 := router.Group("/totp/v1")

	v1.GET("/login", login_handler.Login)

	router.Run(":7230")
}
