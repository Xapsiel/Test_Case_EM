package repository

import (
	"github.com/Xapsiel/EffectiveMobile/internal/models"
	"github.com/jmoiron/sqlx"
)

type Song interface {
	GetSongs(models.Filter) ([]models.Song, error)
	GetSongVerse(song models.Song) (string, int, error)
	DeleteSong(song models.Song) (bool, error)
	UpdateSong(song models.Song) (bool, models.Song, error)
	Add(song models.Song) (bool, int, error)
}
type Repository struct {
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Song: NewSongPostgres(db),
	}
}
