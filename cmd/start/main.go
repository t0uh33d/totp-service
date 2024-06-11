package main

import (
	_ "github.com/lib/pq"
	"github.com/t0uh33d/totp-service/db"
	"github.com/t0uh33d/totp-service/internal/config"
)

func main() {
	// loading the config
	config.LoadConfig()

	db.ConnectToDatabase()

	defer db.SQLx.Close()

	// router := gin.Default()

	// v1 := router.Group("/mbct/v1")

	// router.Run(":7230")
}
