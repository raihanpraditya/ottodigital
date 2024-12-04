package handlers

import (
	"api/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func CreateVoucher(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var voucher models.Voucher

		if err := json.NewDecoder(r.Body).Decode(&voucher); err != nil {
			log.Printf("Invalid JSON payload: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if voucher.BrandID == 0 || voucher.Name == "" || voucher.CostInPoint <= 0 {
			http.Error(w, "BrandID, Name, and CostInPoint are required fields and must be valid", http.StatusBadRequest)
			return
		}

		query := "INSERT INTO vouchers (brand_id, name, cost_in_point) VALUES (?, ?, ?)"
		result, err := db.Exec(query, voucher.BrandID, voucher.Name, voucher.CostInPoint)
		if err != nil {
			log.Printf("Database error during voucher creation: %v", err)
			http.Error(w, "Failed to create voucher", http.StatusInternalServerError)
			return
		}

		lastID, err := result.LastInsertId()
		if err != nil {
			log.Printf("Failed to fetch last inserted ID: %v", err)
			http.Error(w, "Failed to create voucher", http.StatusInternalServerError)
			return
		}
		voucher.ID = int(lastID)

		log.Printf("Voucher created successfully: %+v", voucher)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(voucher)
	}
}

func GetVoucherByID(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		voucherIDStr := r.URL.Query().Get("id")
		if voucherIDStr == "" {
			http.Error(w, "Voucher ID is required", http.StatusBadRequest)
			return
		}

		voucherID, err := strconv.Atoi(voucherIDStr)
		if err != nil {
			http.Error(w, "Invalid voucher ID", http.StatusBadRequest)
			return
		}

		var voucher models.Voucher
		query := "SELECT id, brand_id, name, cost_in_point FROM vouchers WHERE id = ?"
		if err := db.Get(&voucher, query, voucherID); err != nil {
			log.Printf("Database error during voucher retrieval: %v", err)
			http.Error(w, "Voucher not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(voucher)
	}
}

func GetVouchersByBrand(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brandIDStr := r.URL.Query().Get("id")
		if brandIDStr == "" {
			http.Error(w, "Brand ID is required", http.StatusBadRequest)
			return
		}

		brandID, err := strconv.Atoi(brandIDStr)
		if err != nil {
			http.Error(w, "Invalid brand ID", http.StatusBadRequest)
			return
		}

		var vouchers []models.Voucher
		query := "SELECT id, brand_id, name, cost_in_point FROM vouchers WHERE brand_id = ?"
		if err := db.Select(&vouchers, query, brandID); err != nil {
			log.Printf("Database error during vouchers retrieval: %v", err)
			http.Error(w, "Failed to fetch vouchers", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vouchers)
	}
}
