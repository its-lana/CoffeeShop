package model

import "time"

type Cart struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	CustomerID int `gorm:"not null;unique"`
	MerchantID int
	OrderItem  []OrderItem `gorm:"polymorphic:Owner;polymorphicValue:cart;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
