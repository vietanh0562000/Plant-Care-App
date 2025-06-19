package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Name string
}

var instance *Config

var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		godotenv.Load("./config/.env")
		instance = &Config{
			Name: "NotificationService",
		}
	})

	return instance
}

func (c *Config) GetAppPort() string {
	return os.Getenv("APP_PORT")
}

func (c *Config) GetSMTPHost() string {
	return os.Getenv("SMTP_HOST")
}

func (c *Config) GetSMTPPort() int {
	num, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err == nil {
		return 587
	}

	return num
}

func (c *Config) GetSMTPUser() string {
	return os.Getenv("SMTP_USER")
}

func (c *Config) GetSMTPPassword() string {
	return os.Getenv("SMTP_PASSWORD")
}
