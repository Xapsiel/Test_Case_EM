package repository

import (
	"github.com/Xapsiel/EffectiveMobile"
	"github.com/jmoiron/sqlx"
)

type Song interface {
	GetSongs(EffectiveMobile.Filter) ([]EffectiveMobile.Song, error)
	GetSongVerse(song EffectiveMobile.Song) (string, int, error)
	DeleteSong(song EffectiveMobile.Song) (bool, error)
	UpdateSong(song EffectiveMobile.Song) (bool, EffectiveMobile.Song, error)
	Add(song EffectiveMobile.Song) (bool, int, error)
}
type Repository struct {
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Song: NewSongPostgres(db),
	}
}
