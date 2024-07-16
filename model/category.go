package model

import "time"

type Category struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	CategoryName string `gorm:"not null;size:256"`
	MerchantID   int
	Menus        []Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
