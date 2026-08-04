package server

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Prepare sqlx.DB
func NewDB(config *mysql.Config) (*sqlx.DB, error) {
	driverName := "mysql"
	db, err := sqlx.Open(driverName, config.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open failed: %w", err)
	}

	return db, nil
}
