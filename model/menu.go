package model

import "time"

type Menu struct {
	ID                 int     `gorm:"primaryKey;autoIncrement"`
	ProductName        string  `gorm:"not null;size:256"`
	Price              float64 `gorm:"not null"`
	Description        string  `gorm:"size:512"`
	ProductCode        string  `gorm:"not null;unique;size:128"`
	ProductImage       string  `gorm:"size:512"`
	CategoryID         int
	AvailabilityStatus bool `gorm:"not null;default:true"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
