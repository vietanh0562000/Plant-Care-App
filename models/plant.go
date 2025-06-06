package models

import "time"

type Plant struct {
	ID               int       `gorm:"primaryKey"`
	Name             string    `gorm:"not null"`
	ImagePath        string    `gorm:"column:image_url;not null"`
	WateringInterval int       `gorm:"not null"`
	LastTimeWatering time.Time `gorm:"type:timestamp"`
	UserID           int       `gorm:"not null"`
	User             User      `gorm:"foreignKey:UserID"`
	SpeciesID        int       `gorm:"foreignKey:SpeciesID"`
}
