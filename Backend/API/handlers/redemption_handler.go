package handlers

import (
	"api/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func MakeRedemption(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction

		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			log.Printf("Invalid JSON payload: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if transaction.CustomerName == "" || len(transaction.Details) == 0 {
			http.Error(w, "CustomerName and at least one voucher detail are required", http.StatusBadRequest)
			return
		}

		tx, err := db.Begin()
		if err != nil {
			log.Printf("Failed to begin transaction: %v", err)
			http.Error(w, "Failed to process redemption", http.StatusInternalServerError)
			return
		}

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				http.Error(w, "An error occurred while processing the transaction", http.StatusInternalServerError)
			}
		}()

		query := "INSERT INTO transactions (customer_name, total_points) VALUES (?, ?)"
		result, err := tx.Exec(query, transaction.CustomerName, transaction.TotalPoints)
		if err != nil {
			log.Printf("Database error during transaction creation: %v", err)
			tx.Rollback()
			http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
			return
		}

		transactionID, err := result.LastInsertId()
		if err != nil {
			log.Printf("Failed to fetch last transaction ID: %v", err)
			tx.Rollback()
			http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
			return
		}
		transaction.ID = int(transactionID)

		for _, detail := range transaction.Details {
			query = "INSERT INTO transaction_details (transaction_id, voucher_id, quantity) VALUES (?, ?, ?)"
			_, err := tx.Exec(query, transaction.ID, detail.VoucherID, detail.Quantity)
			if err != nil {
				log.Printf("Database error during transaction detail creation: %v", err)
				tx.Rollback()
				http.Error(w, "Failed to process transaction details", http.StatusInternalServerError)
				return
			}
		}

		if err := tx.Commit(); err != nil {
			log.Printf("Failed to commit transaction: %v", err)
			http.Error(w, "Failed to complete redemption", http.StatusInternalServerError)
			return
		}

		log.Printf("Transaction completed successfully: %+v", transaction)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)
	}
}

func GetTransactionDetails(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionIDStr := r.URL.Query().Get("transactionId")
		if transactionIDStr == "" {
			http.Error(w, "Transaction ID is required", http.StatusBadRequest)
			return
		}

		transactionID, err := strconv.Atoi(transactionIDStr)
		if err != nil {
			http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
			return
		}

		var transaction models.Transaction
		query := "SELECT id, customer_name, total_points, created_at FROM transactions WHERE id = ?"
		if err := db.Get(&transaction, query, transactionID); err != nil {
			log.Printf("Database error during transaction retrieval: %v", err)
			http.Error(w, "Transaction not found", http.StatusNotFound)
			return
		}

		var details []models.TransactionDetail
		query = "SELECT id, transaction_id, voucher_id, quantity FROM transaction_details WHERE transaction_id = ?"
		if err := db.Select(&details, query, transactionID); err != nil {
			log.Printf("Database error during transaction details retrieval: %v", err)
			http.Error(w, "Failed to fetch transaction details", http.StatusInternalServerError)
			return
		}
		transaction.Details = details

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)
	}
}
