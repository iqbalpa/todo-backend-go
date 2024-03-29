package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type Config struct {
	// server
	ServerPort string
	// database
	DatabaseDriver string
	DatabaseHost string
	DatabasePort string
	DatabaseName string
	DatabaseUser string
	DatabasePassword string
	// auth
	JwtSecretKey string
	JwtExpiresIn string
}

func NewConfig() *Config {
	// Default value
	return &Config{
		ServerPort: "8080",
		DatabaseDriver: "postgres",
		DatabaseHost: "localhost",
		DatabasePort: "5432",
		DatabaseName: "todo_go",
		DatabaseUser: "postgres",
		DatabasePassword: "",
		JwtSecretKey: "",
		JwtExpiresIn: "",
	}
}

func LoadEnv() *Config {
	config := NewConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading env: \n", err)
	}

	if serverPort := os.Getenv("SERVER_PORT"); serverPort != "" {
		config.ServerPort = serverPort
	}
	if driver := os.Getenv("DATABASE_DRIVER"); driver != "" {
		config.DatabaseDriver = driver
	}
	if host := os.Getenv("DATABASE_HOST"); host != "" {
		config.DatabaseHost = host
	}
	if port := os.Getenv("DATABASE_PORT"); port != "" {
		config.DatabasePort = port
	}
	if name := os.Getenv("DATABASE_NAME"); name != "" {
		config.DatabaseName = name
	}
	if user := os.Getenv("DATABASE_USER"); user != "" {
		config.DatabaseUser = user
	}
	if password := os.Getenv("DATABASE_PASSWORD"); password != "" {
		config.DatabasePassword = password
	}
	if key := os.Getenv("JWT_SECRET_KEY"); key != "" {
		config.JwtSecretKey = key
	}
	if exp := os.Getenv("JWT_EXPIRE_IN_HOur"); exp != "" {
		config.JwtExpiresIn = exp
	}

	return config
}

func (config *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		config.DatabaseHost, 
		config.DatabaseUser, 
		config.DatabasePassword, 
		config.DatabaseName, 
		config.DatabasePort,
	)
}

func (config *Config) GetServerPort() string {
	return fmt.Sprintf("localhost:%s", config.ServerPort)
}

func (config *Config) GetJwtSecretKey() string {
	return config.JwtSecretKey
}

func (config *Config) GetJwtExpiration() string {
	return config.JwtExpiresIn
}