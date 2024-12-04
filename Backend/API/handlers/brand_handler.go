package handlers

import (
	"api/models"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func CreateBrand(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var brand models.Brand

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		log.Printf("Raw Request Body: %s", string(body))

		if err := json.Unmarshal(body, &brand); err != nil {
			log.Printf("Invalid JSON payload: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Printf("Decoded Struct: %+v", brand)

		if brand.Name == "" {
			log.Println("Validation error: 'name' field is required")
			http.Error(w, "The 'name' field is required", http.StatusBadRequest)
			return
		}

		query := "INSERT INTO brands (name) VALUES (?)"
		result, err := db.Exec(query, brand.Name)
		if err != nil {
			log.Printf("Database error during insert: %v", err)
			http.Error(w, "Failed to create brand", http.StatusInternalServerError)
			return
		}

		log.Printf("Insert Result: %+v", result)

		lastID, err := result.LastInsertId()
		if err != nil {
			log.Printf("Error retrieving last inserted ID: %v", err)
			http.Error(w, "Failed to create brand", http.StatusInternalServerError)
			return
		}

		brand.ID = int(lastID)
		log.Printf("Brand created successfully with ID: %d", brand.ID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(brand)
	}
}
