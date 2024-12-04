package models

type Voucher struct {
	ID          int    `json:"id" db:"id"`
	BrandID     int    `json:"brand_id" db:"brand_id"`
	Name        string `json:"name" db:"name"`
	CostInPoint int    `json:"cost_in_point" db:"cost_in_point"`
}
