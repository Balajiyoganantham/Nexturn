package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB1 *sql.DB

// InitializeDatabase1 initializes the products database and creates the products table if not exists
func InitializeDatabase1() error {
	var err error
	DB1, err = sql.Open("sqlite3", "./products.db")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Ping the database to ensure the connection is valid
	if err := DB1.Ping(); err != nil {
		return fmt.Errorf("database connection failed: %v", err)
	}

	// Create Products table if not exists
	_, err = DB1.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		stock INTEGER,
		category TEXT
	);`)
	if err != nil {
		return fmt.Errorf("failed to create products table: %v", err)
	}

	log.Println("Successfully connected to the products database and ensured the schema exists.")
	return nil
}

// GetDB1 returns the products database instance
func GetDB1() *sql.DB {
	if DB1 == nil {
		log.Fatal("Database not initialized. Call InitializeDatabase1 first.")
	}
	return DB1
}
