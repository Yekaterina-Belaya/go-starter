package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    Port        string
    DBUrl       string
    JwtSecret   string
    AppEnv      string
}

var AppConfig *Config

func LoadConfig() {
    // Загружаем .env если он есть
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system env vars")
    }

    AppConfig = &Config{
        Port:      getEnv("PORT", "3000"),
        DBUrl:     getEnv("DATABASE_URL", ""),
        JwtSecret: getEnv("JWT_SECRET", "changeme"),
        AppEnv:    getEnv("APP_ENV", "development"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
