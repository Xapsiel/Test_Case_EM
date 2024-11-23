package repository

import (
	"fmt"
	. "github.com/Xapsiel/EffectiveMobile/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Добавлено для поддержки файлового источника
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	songTable = "song"
	DB_NAME   = "library"
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
		Logger.Debug(MakeLog("Ошибка", err))
		return nil, fmt.Errorf("Ошибка подключения к базе данных postgres")
	}
	Logger.Info("Подключение к базе данных postgres")
	_, err = initialDB.Exec(fmt.Sprintf("CREATE DATABASE %s", DB_NAME))
	if err != nil {
		Logger.Debug("База данных уже создана")
		err = nil
	} else {

		Logger.Debug("База данных создана")
	}
	defer func() {
		initialDB.Close()
		Logger.Info("Отключение от базе данных")
	}()
	mainConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password='%s' sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", mainConnStr)
	if err != nil {
		Logger.Debug(MakeLog("Ошибка", err))
		return nil, fmt.Errorf("Ошибка подключения к базе данных %s", DB_NAME)
	}
	Logger.Info(fmt.Sprintf("Подключение к базе данных %s", DB_NAME))
	if err = db.Ping(); err != nil {
		Logger.Debug(MakeLog("Ошибка", err))
		return nil, fmt.Errorf("Ошибка пингования базы данных")
	}
	Logger.Debug("Проверка подключения прошла успешно")
	err = Migrate(db, cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *sqlx.DB, cfg Config) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		Logger.Debug(MakeLog("Ошибка", err))
		return fmt.Errorf("Ошибка создания объекта драйвера")
	}
	Logger.Debug("Создание объекта драйвера")
	m, err := migrate.NewWithDatabaseInstance(
		"file://schema",
		cfg.DBName, driver,
	)
	if err != nil {
		Logger.Debug(MakeLog("Ошибка", err))
		return fmt.Errorf("Ошибка создания объекта миграция")
	}
	Logger.Debug("Создание объекта миграции")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		Logger.Debug(MakeLog("Ошибка", err))
		return fmt.Errorf("Ошибка принятия миграции")
	}
	Logger.Info("Миграция прошла успешно")
	return nil
}
