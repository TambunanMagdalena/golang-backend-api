package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	PaketDataID uint      `gorm:"not null" json:"paket_data_id"`
	Price       float64   `gorm:"not null" json:"price"`
	CreatedAt   time.Time `json:"created_at"`

	User      User      `gorm:"foreignKey:UserID"`
	PaketData PaketData `gorm:"foreignKey:PaketDataID"`
}
