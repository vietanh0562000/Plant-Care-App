package models

type Species struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Notes string
}
