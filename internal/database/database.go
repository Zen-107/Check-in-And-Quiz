package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB(dsn string) error {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// Migrate creates the necessary tables
func Migrate() error {
	// Create checkins table
	createCheckinsTable := `
	CREATE TABLE IF NOT EXISTS checkins (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_type ENUM('student', 'staff', 'external') NOT NULL,
		name VARCHAR(255) NOT NULL,
		staff_id VARCHAR(100),
		faculty VARCHAR(255),
		position VARCHAR(255),
		occupation VARCHAR(255),
		checkin_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	// Create quiz_results table
	createQuizResultsTable := `
	CREATE TABLE IF NOT EXISTS quiz_results (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		answers JSON NOT NULL,
		result_type VARCHAR(100) NOT NULL,
		result_description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(createCheckinsTable)
	if err != nil {
		return fmt.Errorf("failed to create checkins table: %w", err)
	}

	_, err = DB.Exec(createQuizResultsTable)
	if err != nil {
		return fmt.Errorf("failed to create quiz_results table: %w", err)
	}

	return nil
}
