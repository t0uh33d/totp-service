package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/t0uh33d/totp-service/internal/config"
)

var (
	SQLx *sqlx.DB
)

func ConnectToDatabase() {
	var err error
	SQLx, err = sqlx.Connect("postgres", config.GlobalAppConfig.DB_URL)

	if err != nil {
		log.Fatal(err)
	}

	err = SQLx.Ping()

	if err != nil {
		log.Fatalln("Error: Could not establish a connection with the database", err)
	}

}
