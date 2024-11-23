package repository

import (
	"fmt"
	"github.com/Xapsiel/EffectiveMobile/internal/models"
	. "github.com/Xapsiel/EffectiveMobile/pkg/log"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	pageContain  = 20
	TimeTemplate = "2006-02-02"
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
func (s *SongPostgres) GetSongs(filter models.Filter) ([]models.Song, error) {
	var totalRecords int
	err := s.db.QueryRow("SELECT COUNT(*) FROM songs").Scan(&totalRecords)
	if err != nil {
		return nil, err
	}

	var songs []models.Song
	// Стартовый запрос
	query := "SELECT id,group_name, song_name, release_date, link, text  FROM songs WHERE 1=1"

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

	if filter.Since != "" {
		query += " AND release_date >= $" + fmt.Sprintf("%d", argCount)
		args = append(args, filter.Since)
		argCount++
	}

	if filter.To != "" {
		query += " AND release_date <= $" + fmt.Sprintf("%d", argCount)
		args = append(args, filter.To)
		argCount++
	}
	if filter.ID != 0 {
		query += " AND id = $" + fmt.Sprintf("%d", argCount)
		args = append(args, filter.ID)
		argCount++
	}
	query += fmt.Sprintf(" ORDER BY id DESC LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, pageContain)
	args = append(args, ((filter.Page - 1) * pageContain))
	// Выполняем запрос
	err = s.db.Select(&songs, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения списка песен")
	}
	Logger.Debug(fmt.Sprintf("Запрос к базе данных:%s с аргументами [%v]", query, args))

	return songs, nil
}

func (s *SongPostgres) GetSongVerse(song models.Song) (string, int, error) {
	var result []struct {
		Text string `json: "text"`
		Id   int    `json: "id"`
	}
	query := "SELECT text, id FROM songs WHERE "
	if song.SongName != "" {
		query += "(song_name = $1 AND group_name = $2)"
		err := s.db.Select(&result, query, song.SongName, song.Group)
		if err != nil {
			return "", 0, fmt.Errorf("Ошибка получения песни")
		}
		Logger.Debug(fmt.Sprintf("Запрос к базе данных:%s с аргументами [%v]", query, song.SongName, song.Group))
		return result[0].Text, result[0].Id, nil
	} else if song.ID != 0 {
		query += "id  = $1"

		err := s.db.Get(&result, query, song.ID)
		if err != nil {
			return "", 0, fmt.Errorf("Ошибка получения песни")
		}
		Logger.Debug(fmt.Sprintf("Запрос к базе данных:%s с аргументами [%v]", query, song.ID))
		return result[0].Text, result[0].Id, nil
	}
	return "", 0, fmt.Errorf("Песня не найдена")

}
func (s *SongPostgres) DeleteSong(song models.Song) (bool, error) {
	query := "DELETE FROM songs WHERE "
	if song.SongName != "" {
		query += "(song_name = $1 AND group_name = $2)"
		_, err := s.db.Exec(query, song.SongName, song.Group)
		if err != nil {
			return false, fmt.Errorf("Ошибка удаления песни")
		}
		return true, nil
	} else if song.ID != 0 {
		query += "id  = $1"
		_, err := s.db.Exec(query, song.ID)
		if err != nil {
			return false, fmt.Errorf("Ошибка удаления песни")
		}
		return true, nil
	}
	return false, fmt.Errorf("Песня не найдена")

}
func (s *SongPostgres) UpdateSong(song models.Song) (bool, models.Song, error) {
	query := "UPDATE songs SET id=id "
	var args []interface{}
	argCount := 1
	realesedDate, err := time.Parse(TimeTemplate, song.ReleaseDate)

	if err != nil && song.ReleaseDate != "" {
		Logger.Debug(MakeLog("Ошибка", err))
		return false, models.Song{}, fmt.Errorf("Не тот формат даты")
	} else if song.ReleaseDate == "" {
		realesedDate = time.Time{}
	}
	// Проверяем каждый параметр фильтра и добавляем его в запрос, если он установлен
	if song.Text != "" {
		query += " , text = $" + fmt.Sprintf("%d", argCount)
		args = append(args, song.Text)
		argCount++
	}
	if song.Link != "" {
		query += " , link = $" + fmt.Sprintf("%d", argCount)
		args = append(args, song.Link)
		argCount++
	}
	if !realesedDate.IsZero() {
		query += " , release_date = $" + fmt.Sprintf("%d", argCount)
		args = append(args, song.ReleaseDate)
		argCount++
	}
	if song.SongName != "" {
		query += fmt.Sprintf(" WHERE song_name = $%d AND group_name = $%d", argCount, argCount+1)
		args = append(args, song.SongName)
		args = append(args, song.Group)
	} else if song.ID != 0 {
		query += fmt.Sprintf(" WHERE id = $%d", argCount)
		args = append(args, song.ID)

	}
	_, err = s.db.Exec(query, args...)
	if err != nil {
		Logger.Debug(MakeLog("Ошибка", err))
		return false, models.Song{}, fmt.Errorf("Ошибка обновления данных о песне")
	}
	Logger.Debug(fmt.Sprintf("Запрос к базе данных:%s с аргументами [%v]", query, args))

	return true, song, nil
}
func (s *SongPostgres) Add(song models.Song) (bool, int, error) {
	query := "INSERT INTO songs (group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := s.db.QueryRow(query, song.Group, song.SongName, song.ReleaseDate, song.Text, song.Link).Scan(&id)
	if err != nil {
		Logger.Debug(MakeLog("Ошибка", err))
		return false, 0, fmt.Errorf("Ошибка добавления песни")
	}
	Logger.Debug(fmt.Sprintf("Запрос к базе данных:%s с аргументами [%v]", query, []interface{}{song.Group, song.SongName, song.ReleaseDate, song.Text, song.Link}))
	return true, id, nil

}
