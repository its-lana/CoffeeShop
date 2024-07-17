package model

import "time"

type Payment struct {
	ID            int       `gorm:"primaryKey;autoIncrement"`
	CustomerID    int       `gorm:"not null"`
	OrderUID      string    `gorm:"not null;size:64"`
	PaymentAmount int       `gorm:"not null"`
	Status        string    `gorm:"not null;size:64"`
	PaymentMethod string    `gorm:"not null;size:64"`
	PaymentURL    string    `gorm:"size:512"`
	PaidDate      time.Time `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
