package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Name string
}

// Biến lưu instance singleton
var instance *Config

// sync.Once để đảm bảo chỉ gọi init một lần
var once sync.Once

// GetInstance trả về instance của Config
func GetInstance() *Config {
	once.Do(func() {
		fmt.Println("Initializing Config...")
		godotenv.Load("./config/.env")
		instance = &Config{Name: "PlantService"}
	})
	return instance
}

func (cfg *Config) GetAppPort() string {
	host := os.Getenv("APP_PORT")
	if host == "" {
		host = "8080"
	}

	return host
}

func (cfg *Config) GetDBHost() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	return host
}

func (cfg *Config) GetDBPort() string {
	host := os.Getenv("DB_PORT")
	if host == "" {
		host = "5432"
	}

	return host
}

func (cfg *Config) GetDBName() string {
	return os.Getenv("DB_NAME")
}

func (cfg *Config) GetDBUser() string {
	return os.Getenv("DB_USER")
}

func (cfg *Config) GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func (cfg *Config) GetUserServiceHost() string {
	return os.Getenv("USER_SERVICE")
}

func (cgf *Config) GetUploadDir() string {
	fmt.Println("GET UPLOAD DIR----------")
	return os.Getenv("UPLOAD_PLANT_DIR")
}
