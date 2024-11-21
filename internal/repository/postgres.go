package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Добавлено для поддержки файлового источника
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	songTable = "song"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// Подключение к базе "postgres" для проверки и создания основной БД
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password='%s' sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode)
	initialDB, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to initial database: %w", err)
	}
	defer initialDB.Close()
	mainConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password='%s' sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", mainConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open main database connection: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping main database: %w", err)
	}

	// Настраиваем миграции
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://schema", // Убедитесь, что путь правильный
		cfg.DBName, driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	return db, nil
}
