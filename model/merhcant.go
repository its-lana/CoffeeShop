package model

import (
	"time"
)

type Merchant struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	MerchantName string `gorm:"not null;size:256"`
	Address      string `gorm:"not null;size:512"`
	PICName      string `gorm:"not null;size:256"`
	Email        string `gorm:"unique;not null;size:64"`
	Password     string `gorm:"not null;size:128"`
	PhoneNumber  string `gorm:"not null;size:16"`
	IsOpen       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
