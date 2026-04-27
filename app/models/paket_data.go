package models

import "time"

type PaketData struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Price        float64   `gorm:"not null" json:"price"`
	Quota        float64   `gorm:"not null" json:"quota"`
	ActivePeriod int       `gorm:"not null" json:"active_period"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Transactions []Transaction `gorm:"foreignKey:PaketDataID"`
}

