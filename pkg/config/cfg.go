package config

import (
	"effective-mobile-task/pkg/logger"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbaseCfg struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

// Загрузка данных из ENV
func EnvLoad(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Log.Fatalf("env variable %s is not set", key)
	}
	return value
}

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		logger.LogError("Error loading .env file", err)
	}
}

// Собираю конфиг
func Config() *DbaseCfg {
	return &DbaseCfg{
		DBUser: EnvLoad("DB_USER"),
		DBPass: EnvLoad("DB_PASSWORD"),
		DBHost: EnvLoad("DB_HOST"),
		DBPort: EnvLoad("DB_PORT"),
		DBName: EnvLoad("DB_NAME"),
	}
}

// URL
func (c *DbaseCfg) DBaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}
