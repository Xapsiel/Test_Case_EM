package service

import (
	"github.com/Xapsiel/EffectiveMobile/internal/models"
	"github.com/Xapsiel/EffectiveMobile/internal/repository"
	"io"
)

type Song interface {
	GetSongs(filter models.Filter) ([]models.Song, error)
	GetSongVerse(song models.Song, verse int) (string, int, error)
	DeleteSong(song models.Song) (bool, error)
	UpdateSong(song models.Song) (bool, models.Song, error)
	Add(song models.Song) (bool, int, error)
}
type Log interface {
	SetFormat(io.Writer)
	SetLevel(lvl string)
	Info(string)
	Warn(string)
	Debug(string)
}
type Service struct {
	Song
	Log
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Song: NewSongService(repo.Song),
	}
}
