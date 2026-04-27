package models

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	PhoneNumber string    `gorm:"type:varchar(20);not null;uniqueIndex" json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Transactions []Transaction `gorm:"foreignKey:PaketDataID" json:"-"`
}
