package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
	ServerPort string
}

func Load() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "jcuadrado"),
		DBPass:     getEnv("DB_PASS", "jcuadrado"),
		DBName:     getEnv("DB_NAME", "go_lab"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
