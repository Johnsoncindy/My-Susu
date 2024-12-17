package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Johnsoncindy/My-Susu/my-susu-backend/internal/database"
)

type APIConfig struct {
	DB         *database.Queries
	DBConn     *sql.DB
	Logger     *log.Logger
	MoMoConfig *MoMoConfig
}

func InitConfig() (*APIConfig, error) {
	// Load config from environment variables
	dbURL := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Initialize queries
	queries := database.New(conn)

	// Initialize logger
	logger := log.New(os.Stdout, "[My-Susu]", log.LstdFlags|log.Lshortfile)

	// Initialize MoMoConfig
	moMoConfig, err := LoadMoMoConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load MoMo config: %v", err)
	}

	return &APIConfig{
		DB:         queries,
		DBConn:     conn,
		Logger:     logger,
		MoMoConfig: moMoConfig,
	}, nil
}
