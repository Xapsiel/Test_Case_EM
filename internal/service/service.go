package service

import (
	"github.com/Xapsiel/EffectiveMobile"
	"github.com/Xapsiel/EffectiveMobile/internal/repository"
)

type Song interface {
	GetSongs(filter EffectiveMobile.Filter) ([]EffectiveMobile.Song, error)
	GetSongVerse(song EffectiveMobile.Song, verse int) (string, int, error)
	DeleteSong(song EffectiveMobile.Song) (bool, error)
	UpdateSong(song EffectiveMobile.Song) (bool, EffectiveMobile.Song, error)
	Add(song EffectiveMobile.Song) (bool, int, error)
}
type Service struct {
	Song
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Song: NewSongService(repo.Song),
	}
}
