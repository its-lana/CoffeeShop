package model

import (
	"time"
)

type OrderItem struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	MenuID    int    `gorm:"not null"`
	Quantity  int    `gorm:"not null"`
	OwnerID   int    `gorm:"not null"`
	OwnerType string `gorm:"not null;size:64"`
	Menu      Menu   `gorm:"foreignKey:MenuID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
