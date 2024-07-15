package model

import "time"

type Customer struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	FullName    string `gorm:"not null;size:256"`
	Email       string `gorm:"unique;not null;size:64"`
	PhoneNumber string `gorm:"size:16"`
	Password    string `gorm:"size:128"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
