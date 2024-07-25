package config

import (
	"fmt"
	"os"
	"path/filepath"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress  string
	ImmudbURL      string
	ImmudbUsername string
	ImmudbPassword string
}

// Read Environment variables from eoth .env file or system 
func Load() (*Config, error) {
	projectRoot := filepath.Join("..", "..") 
	envFilePath := filepath.Join(projectRoot, ".env")
	log.Println("Loading .env file from:", envFilePath)

	if err := godotenv.Load(envFilePath); err != nil {
		log.Println("Warning: .env file not found. Using environment variables from system.")
	}

	config := &Config{
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
		ImmudbUsername: os.Getenv("IMMUDB_USERNAME"),
		ImmudbPassword: os.Getenv("IMMUDB_PASSWORD"),
	}

	if config.ServerAddress == "" || config.ImmudbUsername == "" || config.ImmudbPassword == "" {
		return nil, fmt.Errorf("missing required immudb related values or  ServerAddress")
	}

	return config, nil
}
