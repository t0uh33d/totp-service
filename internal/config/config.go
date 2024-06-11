package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_URL string
}

var GlobalAppConfig *AppConfig

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// OPEN_AI_KEY := extractEnv("OPENAI_API_KEY")

	// // Only initialize GlobalAppConfig if both required configurations are provided.
	// if OPEN_AI_KEY != "" {
	// 	// bot_handler.Client = openai.NewClient(OPEN_AI_KEY)
	// 	GlobalAppConfig = &AppConfig{
	// 		DB_URL: "",
	// 	}
	// } else {
	// 	fmt.Println("Failed to load configuration: Missing required environment variables.")
	// 	os.Exit(1) // Exit the program if critical configuration is missing
	// }
}

func extractEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Failed to extract %s from environment variables.\n", key)
	}
	return value
}
