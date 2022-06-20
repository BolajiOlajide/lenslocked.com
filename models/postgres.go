package models

import (
	"database/sql"
	"fmt"

	// adding this here to initialize the pg driver in pgx
	_ "github.com/jackc/pgx/v4/stdlib"
)

// PostgresConfig configuration information for database connection
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// ToString returns the postgres url string using the credentials passed
func (cfg PostgresConfig) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)
}

// DefaultPostgresConfig returns the default config for connecting to PG in development
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

// Open will open a SQL connection with the provided postgres database.
// Callers of Open need to ensure that the connection is eventually closed
// via the db.Close() method.
func Open(cfg PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		return nil, fmt.Errorf("DB open: %w", err)
	}

	return db, nil
}
