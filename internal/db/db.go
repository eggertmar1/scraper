package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// NewDB creates a new database connection using sqlx
func NewDB() (*sqlx.DB, error) {
	//Database connection settings
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open connection to PostgreSQL database
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Verify the connection is working
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot reach the database:", err)
		return nil, err
	}

	log.Println("Connected to the database successfully.")
	return db, nil
}
