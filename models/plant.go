package models

type Plant struct {
	ID               int    `gorm:"primaryKey"`
	Name             string `gorm:"not null"`
	ImageURL         string `gorm:"not null"`
	WateringInterval int    `gorm:"not null"`
	UserID           int    `gorm:"not null"`
	User             User   `gorm:"foreignKey:UserID"`
}
