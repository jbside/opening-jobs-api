package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeDB() (*sqlx.DB, error) {
	var err error
	var resultDB *sqlx.DB

	// Create database connection
	resultDB, err = sqlx.Connect("postgres", GetEnv("URL_DB", ""))
	if err != nil {
		err := fmt.Errorf("database connection error: %v", err)
		return nil, err
	}

	// Create schema app
	_, err = resultDB.Exec(`CREATE SCHEMA IF NOT EXISTS app;`)
	if err != nil {
		err := fmt.Errorf("create database extension error: %v", err)
		return nil, err
	}

	return resultDB, nil
}
