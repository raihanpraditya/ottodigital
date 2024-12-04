package models

import "time"

type Transaction struct {
	ID           int                 `json:"id" db:"id"`
	CustomerName string              `json:"customer_name" db:"customer_name"`
	TotalPoints  int                 `json:"total_points" db:"total_points"`
	CreatedAt    time.Time           `json:"created_at" db:"created_at"`
	Details      []TransactionDetail `json:"details,omitempty"`
}

type TransactionDetail struct {
	ID            int `json:"id" db:"id"`
	TransactionID int `json:"transaction_id" db:"transaction_id"`
	VoucherID     int `json:"voucher_id" db:"voucher_id"`
	Quantity      int `json:"quantity" db:"quantity"`
}
