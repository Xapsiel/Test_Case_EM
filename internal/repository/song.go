package repository

import (
	"fmt"
	"github.com/Xapsiel/EffectiveMobile"
	"github.com/jmoiron/sqlx"
)

type SongPostgres struct {
	db *sqlx.DB
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

/*
CREATE TABLE songs (

	id SERIAL PRIMARY KEY,
	group_name VARCHAR(255) NOT NULL,
	song_name VARCHAR(255) NOT NULL,
	release_date DATE NOT NULL,
	text TEXT NOT NULL,
	link VARCHAR(255) NOT NULL

);
*/
func (s *SongPostgres) GetSongs(filter EffectiveMobile.Filter) ([]EffectiveMobile.Song, error) {
	var songs []EffectiveMobile.Song

	// Стартовый запрос
	query := "SELECT release_date, text, link FROM songs WHERE 1=1"

	// Переменная для хранения параметров запроса
	var args []interface{}
	argCount := 1

	// Проверяем каждый параметр фильтра и добавляем его в запрос, если он установлен
	if filter.Group != "" {
		query += " AND group_name = $" + fmt.Sprintf("%d", argCount)
		args = append(args, filter.Group)
		argCount++
	}

	if filter.Song != "" {
		query += " AND song_name = $" + fmt.Sprintf("%d", argCount)
		args = append(args, filter.Song)
		argCount++
	}

	if !filter.Since.IsZero() {
		query += " AND release_date >= $" + fmt.Sprintf("%d", argCount)
		args = append(args, filter.Since)
		argCount++
	}

	// Проверяем, есть ли значение для поля EndDate (time.Time)
	if !filter.To.IsZero() {
		query += " AND release_date <= $" + fmt.Sprintf("%d", argCount)
		args = append(args, filter.To)
		argCount++
	}

	// Выполняем запрос
	err := s.db.Select(&songs, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения списка песен")
	}

	return songs, nil
}

func (s *SongPostgres) GetSongVerse(song, group string) (string, error) {
	var text string
	query := "SELECT text FROM songs WHERE song_name=$1 AND group_name = $2"
	err := s.db.Select(&text, query, song, group)
	if err != nil {
		return "", fmt.Errorf("Ошибка получения песни")
	}
	return text, nil
}
func (s *SongPostgres) DeleteSong(song string, group string) (bool, error) {
	query := "DELETE FROM songs WHERE song_name=$1 AND group_name = $2"
	_, err := s.db.Exec(query, song, group)
	if err != nil {
		return false, fmt.Errorf("Ошибка удаления песни")
	}
	return true, nil
}
func (s *SongPostgres) UpdateSong(song EffectiveMobile.Song) (bool, error) {
	query := "UPDATE songs SET id=id "
	var args []interface{}
	argCount := 1

	// Проверяем каждый параметр фильтра и добавляем его в запрос, если он установлен
	if song.Text != nil {
		query += " AND text = $" + fmt.Sprintf("%d", argCount)
		args = append(args, *song.Text)
		argCount++
	}
	if song.Link != nil {
		query += " AND ling = $" + fmt.Sprintf("%d", argCount)
		args = append(args, *song.Link)
		argCount++
	}
	if !song.ReleaseDate.IsZero() {
		query += " AND release_date = $" + fmt.Sprintf("%d", argCount)
		args = append(args, song.ReleaseDate)
		argCount++
	}
	query += fmt.Sprintf("WHERE song_name = $%d AND group_name = $%d", argCount, argCount+1)
	_, err := s.db.Exec(query, args)
	if err != nil {
		return false, fmt.Errorf("Ошибка обновления данных о песне")
	}

	return true, nil
}
func (s *SongPostgres) Add(song EffectiveMobile.Song) (bool, error) {
	query := "INSERT INTO songs (group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5)"

	_, err := s.db.Exec(query, song.Group, song.SongName, song.ReleaseDate, *song.Text, *song.Link)
	if err != nil {
		return false, fmt.Errorf("Ошибка добавления песни")
	}
	return true, nil

}
