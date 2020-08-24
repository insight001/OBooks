package config

import "os"

//DatabaseConfig ...
type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	Password string
	URL      string
	User     string
}

//Config ...
type Config struct {
	Database DatabaseConfig
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			Name:     getEnv("DB_NAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
			URL:      getEnv("DB_URL", ""),
			User:     getEnv("DB_USER", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
