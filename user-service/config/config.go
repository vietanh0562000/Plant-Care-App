package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost       string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassword    string
	DbName        string
	JwtSecret     string
	AdminEmail    string
	AdminPassword string
	AdminName     string
}

func LoadConfig() *Config {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Println("⚠️ Không tìm thấy file .env, sẽ dùng biến môi trường hiện tại")
	}

	return &Config{
		AppHost:       getEnv("APP_PORT", "8080"),
		DbHost:        getEnv("DB_HOST", "localhost"),
		DbPort:        getEnv("DB_PORT", "5432"),
		DbUser:        getEnv("DB_USER", "postgres"),
		DbPassword:    getEnv("DB_PASSWORD", ""),
		DbName:        getEnv("DB_NAME", ""),
		JwtSecret:     getEnv("JWT_SECRET", "secret"),
		AdminEmail:    getEnv("ADMIN_EMAIL", ""),
		AdminPassword: getEnv("ADMIN_PASSWORD", ""),
		AdminName:     getEnv("ADMIN_NAME", ""),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
