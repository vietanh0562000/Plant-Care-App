package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Time    time.Time
	PlantID uint `gorm:"foreignKey:PlantID"`
	Note    string
}
