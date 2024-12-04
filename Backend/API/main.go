package main

import (
	"log"
	"net/http"

	"api/database"
	"api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to apply migrations:", err)
	}
	r := mux.NewRouter()

	// Brand Routes
	r.HandleFunc("/brand", handlers.CreateBrand(db)).Methods("POST")
	// Voucher Routes
	r.HandleFunc("/voucher", handlers.CreateVoucher(db)).Methods("POST")
	r.HandleFunc("/voucher", handlers.GetVoucherByID(db)).Methods("GET")
	r.HandleFunc("/voucher/brand", handlers.GetVouchersByBrand(db)).Methods("GET")
	// Transaction Routes
	r.HandleFunc("/transaction/redemption", handlers.MakeRedemption(db)).Methods("POST")
	r.HandleFunc("/transaction/redemption", handlers.GetTransactionDetails(db)).Methods("GET")

	// Start Server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
