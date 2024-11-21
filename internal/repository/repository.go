package repository

import (
	"github.com/Xapsiel/EffectiveMobile"
	"github.com/jmoiron/sqlx"
)

type Song interface {
	GetSongs(EffectiveMobile.Filter) ([]EffectiveMobile.Song, error)
	GetSongVerse(string, string) (string, error)
	DeleteSong(string, string) (bool, error)
	UpdateSong(EffectiveMobile.Song) (bool, error)
	Add(song EffectiveMobile.Song) (bool, error)
}
type Repository struct {
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Song: NewSongPostgres(db),
	}
}
