package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/pressly/goose"
	"github.com/rhajizada/donezo-mini/internal/tui/app"

	"github.com/rhajizada/donezo-mini/internal/repository"
	"github.com/rhajizada/donezo-mini/internal/service"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicf("unable to determine user home directory: %v", err)
	}
	donezoDir := filepath.Join(homeDir, ".donezo")
	_, err = os.Stat(donezoDir)
	if os.IsNotExist(err) {
		os.Mkdir(donezoDir, 0700)
	}

	dbPath := filepath.Join(donezoDir, "data.db")

	// Load database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Panicf("failed to open database %s: %v", dbPath, err)
	}
	defer db.Close()

	// Ensure the migrations directory exists
	migrationsDir := "data/sql/migrations"
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		log.Panicf("migrations directory does not exist: %s", migrationsDir)
	}

	// Set Goose dialect to SQLite
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Panicf("failed to set Goose dialect: %v", err)
	}

	// Apply all up migrations
	if err := goose.Up(db, migrationsDir); err != nil {
		log.Panicf("failed to apply migrations: %v", err)
	}

	r := repository.New(db)
	s := service.New(r)
	ctx := context.Background()

	m := app.NewModel(ctx, s)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Panicf("error running program: %v", err)
	}
}
