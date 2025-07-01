package migration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	migrationsDir := filepath.Join(rootPath, "db", "migration", "migration_files")

	migrationsDir = strings.ReplaceAll(migrationsDir, "\\", "/")

	migrationsPath := fmt.Sprintf("file://%s", migrationsDir)

	m, err := migrate.New(
		migrationsPath,
		"postgres://postgres:postgres@localhost:5433/emailn?sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Error initializing migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error when applying migrations: %v", err)
	}

	log.Println("Successfully applied migrations")
}
