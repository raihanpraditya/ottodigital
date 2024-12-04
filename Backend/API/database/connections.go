package database

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/voucher_api?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connected to the database successfully")
	return db, nil
}
