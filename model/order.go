package model

import "time"

type Order struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	OrderUID    string      `gorm:"not null;unique;size:64"`
	CustomerID  int         `gorm:"not null"`
	MerchantID  int         `gorm:"not null"`
	FinalAmount int         `gorm:"not null"`
	OrderType   string      `gorm:"not null;size:64"`
	OrderNotes  string      `gorm:"size:512"`
	OrderStatus string      `gorm:"not null;size:64"`
	NoteStatus  string      `gorm:"size:256"`
	OrderDate   time.Time   `gorm:"not null"`
	OrderCode   string      `gorm:"not null;size:8"`
	Payment     Payment     `gorm:"foreignKey:OrderUID;references:OrderUID"`
	OrderItem   []OrderItem `gorm:"polymorphic:Owner;polymorphicValue:order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
