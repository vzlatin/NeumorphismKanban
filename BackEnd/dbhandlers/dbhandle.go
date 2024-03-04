package dbhandlers

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/vzlatin/NeumorphismKanban/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

var apiConfig *ApiConfig = NewApiConfig(getDBURL())

func NewApiConfig(DB_URL string) *ApiConfig {
	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Printf("ERROR: Can't connect to the database: %s", err)
	}

	queries := database.New(conn)

	return &ApiConfig{
		DB: queries,
	}
}

func getDBURL() string {
	// Load the .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %s", err)
	}

	// Get the DB_URL from the environment
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL not found in .env file or environment variables")
	}

	return dbURL
}

func GetApiConfig() (*ApiConfig, context.Context) {
	ctx := context.Background()
	return apiConfig, ctx
}
