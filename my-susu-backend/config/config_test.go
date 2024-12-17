package config

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock the database.Queries for testing
type MockQueries struct {
	mock.Mock
}

func TestInitConfig(t *testing.T) {
	// Set up environment variable for the test
	os.Setenv("DB_URL", "postgres://default:g8mNio0ubBVK@ep-long-fog-a48agn76.us-east-1.aws.neon.tech:5432/verceldb?sslmode=require")
	defer os.Unsetenv("DB_URL") // Clean up environment variable after test

	// Mock sql.Open and sql.DB.Ping
	originalSqlOpen := sqlOpen
	originalDbPing := dbPing

	defer func() {
		sqlOpen = originalSqlOpen
		dbPing = originalDbPing
	}()

	sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		return &sql.DB{}, nil
	}
	dbPing = func(db *sql.DB) error {
		return nil
	}

	// Test InitConfig function
	config, err := InitConfig()
	require.NoError(t, err, "InitConfig should not return an error")
	assert.NotNil(t, config, "config should not be nil")
	assert.NotNil(t, config.DB, "config.DB should not be nil")
}

var sqlOpen = sql.Open
var dbPing = (*sql.DB).Ping
