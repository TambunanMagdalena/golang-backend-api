package models

// =========================
// USER REQUEST
// =========================

type CreateUserRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateUserRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

// =========================
// PAKET DATA REQUEST
// =========================

type CreatePaketDataRequest struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Quota        float64 `json:"quota"`
	ActivePeriod int     `json:"active_period"`
}

type UpdatePaketDataRequest struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Quota        float64 `json:"quota"`
	ActivePeriod int     `json:"active_period"`
}

// =========================
// TRANSACTION REQUEST
// =========================

type CreateTransactionRequest struct {
	UserID      uint `json:"user_id"`
	PaketDataID uint `json:"paket_data_id"`
}