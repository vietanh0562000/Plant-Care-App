package models

type Plant struct {
	ID               int    `gorm:"primaryKey"`
	Name             string `gorm:"not null"`
	ImageURL         string `gorm:"not null"`
	WateringInterval int    `gorm:"not null"`
}
