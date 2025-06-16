package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	IsAdmin   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
