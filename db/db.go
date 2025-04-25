package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect to PostgreSQL database using pgxpool
var DB *pgxpool.Pool

func Connect() error {
	// Construct the PostgreSQL connection URL using environment variables
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	// trying to connect to the database with retries up to 10 times
	// with a 2 second delay between each attempt
	for i := 0; i < 10; i++ {
		// Attempt to connect to the database using pgx driver
		DB, err = pgxpool.New(context.Background(), url)
		if err == nil {
			err = DB.Ping(context.Background()) // Check if the connection is alive
			if err == nil {
				fmt.Println("Connected to PostgreSQL!")
				return nil // Connection successful
			}
		}
		// Log and wait before retrying
		fmt.Println("Retrying DB connection...")
		time.Sleep(2 * time.Second)
	}
	// Return error if connection could not be established after all retries
	return fmt.Errorf("failed to connect after retries: %w", err)
}
