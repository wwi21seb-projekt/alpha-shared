package config

import (
	"fmt"
	"os"
)

// AlphaConfig holds all possible configuration values
type AlphaConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	SchemaName       string
	Port             string
	UserServiceURL   string
	PostServiceURL   string
}

// LoadConfig loads environment variables into an AlphaConfig struct
func LoadConfig() (*AlphaConfig, error) {
	config := &AlphaConfig{
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresDB:       getEnv("POSTGRES_DB", "mydatabase"),
		PostgresUser:     getEnv("POSTGRES_USER", "myuser"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "mypassword"),
		SchemaName:       getEnv("SCHEMA_NAME", "public"),
		Port:             getEnv("PORT", "8080"),
		UserServiceURL:   getEnv("USER_SERVICE_URL", "http://localhost:50051"),
		PostServiceURL:   getEnv("POST_SERVICE_URL", "http://localhost:50052"),
	}

	// Add any necessary validation here
	if config.Port == "" {
		return nil, fmt.Errorf("missing required environment variable: PORT")
	}

	return config, nil
}

// getEnv reads an environment variable and returns its value or a default value if not set
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}