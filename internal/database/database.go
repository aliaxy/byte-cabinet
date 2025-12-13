package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

// DB wraps sqlx.DB to provide additional functionality
type DB struct {
	*sqlx.DB
	migrationPath string
}

// New creates a new database connection
func New(dbPath string) (*DB, error) {
	// Ensure the directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Open database connection
	// Using modernc.org/sqlite which is a pure Go SQLite implementation
	db, err := sqlx.Connect("sqlite", dbPath+"?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(1) // SQLite only supports one writer at a time
	db.SetMaxIdleConns(1)

	return &DB{
		DB:            db,
		migrationPath: "file://migrations",
	}, nil
}

// NewWithMigrationPath creates a new database connection with custom migration path
func NewWithMigrationPath(dbPath, migrationPath string) (*DB, error) {
	db, err := New(dbPath)
	if err != nil {
		return nil, err
	}
	db.migrationPath = migrationPath
	return db, nil
}

// Migrate runs database migrations
func (db *DB) Migrate() error {
	// Create migration driver
	driver, err := sqlite.WithInstance(db.DB.DB, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	// Create migrator
	m, err := migrate.NewWithDatabaseInstance(db.migrationPath, "sqlite", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

// MigrateDown rolls back all migrations
func (db *DB) MigrateDown() error {
	driver, err := sqlite.WithInstance(db.DB.DB, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(db.migrationPath, "sqlite", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migrations: %w", err)
	}

	return nil
}

// MigrateSteps runs n migration steps (positive = up, negative = down)
func (db *DB) MigrateSteps(n int) error {
	driver, err := sqlite.WithInstance(db.DB.DB, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(db.migrationPath, "sqlite", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	if err := m.Steps(n); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migration steps: %w", err)
	}

	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}

// HealthCheck verifies the database connection is working
func (db *DB) HealthCheck() error {
	return db.DB.Ping()
}
