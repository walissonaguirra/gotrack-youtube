package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// DB encapsula a conexão sql.DB.
type DB struct {
	conn *sql.DB
}

// New abre (ou cria) o banco de dados SQLite no diretório de configuração do usuário.
func New() (*DB, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "."
	}

	dbDir := filepath.Join(dir, "gotrack")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		return nil, fmt.Errorf("erro ao criar diretório do banco: %w", err)
	}

	dbPath := filepath.Join(dbDir, "gotrack.db")
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir banco de dados: %w", err)
	}

	db := &DB{conn: conn}
	if err := db.migrate(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("erro ao executar migrações: %w", err)
	}
	return db, nil
}

// Close fecha a conexão com o banco de dados.
func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) migrate() error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS lessons (
			id TEXT PRIMARY KEY,
			chapter INTEGER NOT NULL,
			lesson_number INTEGER NOT NULL,
			title TEXT NOT NULL,
			youtube_id TEXT NOT NULL,
			is_exercise INTEGER DEFAULT 0,
			completed INTEGER DEFAULT 0,
			completed_at TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS timer_sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			started_at TEXT NOT NULL,
			duration_minutes INTEGER NOT NULL,
			chapter INTEGER
		)`,
	}

	for _, m := range migrations {
		if _, err := db.conn.Exec(m); err != nil {
			return err
		}
	}

	return nil
}
