package config

import (
	"fmt"
	"os"
	"path/filepath"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  	string
	ImmuDbUrl 		string
	ImmuDbApiKey 	string
}

// Read Environment variables from first .env file and if not present then from system 
func Load() (*Config, error) {
	projectRoot := filepath.Join("..", "..") 
	envFilePath := filepath.Join(projectRoot, ".env")
	log.Println("Loading .env file from:", envFilePath)

	if err := godotenv.Load(envFilePath); err != nil {
		log.Println("Warning: .env file not found. Using environment variables from system.")
	}

	config := &Config{
		ServerPort:  	os.Getenv("SERVER_PORT"),
		ImmuDbUrl: 		os.Getenv("IMMUDB_URL"),
		ImmuDbApiKey: 	os.Getenv("IMMUDB_API_KEY"),
	}

	if config.ServerPort == ""  || config.ImmuDbUrl == "" || config.ImmuDbApiKey == "" {
		return nil, fmt.Errorf("missing ServerPort or ImmuDbUrl or ImmuDbApiKey")
	}

	return config, nil
}
