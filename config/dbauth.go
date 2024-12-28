package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// Global variable to hold the database connection for users
var UsersDB *sql.DB

// InitializeDatabase_auth initializes the users database for authentication
func InitializeDatabase_auth() (*sql.DB, error) {
	// Open the connection to the users database (users.db)
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}

	// Create the users table if it does not exist
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	// Store the database connection globally for use in the app
	UsersDB = db

	// Ensure that the users table exists
	fmt.Println("Connected to the users database and ensured the users table exists.")
	return db, nil
}

// GetDB returns the global users database connection
func GetDB2() *sql.DB {
	return UsersDB // Return the connection to the users database
}
