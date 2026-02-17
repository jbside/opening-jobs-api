package config

import (
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	db     *sqlx.DB
	qb     *squirrel.StatementBuilderType
	logger *Logger
)

func Init() error {
	var err error

	// Initialize environment variables from .env file
	_ = godotenv.Load()

	db, err = InitializeDB()
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	qb = InitializeQueryBuilder()

	return nil
}

func GetDB() *sqlx.DB {
	return db
}

func GetQueryBuilder() *squirrel.StatementBuilderType {
	return qb
}

func GetLogger() *Logger {
	if logger == nil {
		logger = NewLogger()
	}
	return logger
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
