package service

import (
	"github.com/Xapsiel/EffectiveMobile"
)

type Song interface {
	GetSongs(EffectiveMobile.Filter) ([]EffectiveMobile.Song, error)
	GetSongVerse(string, string, int) (string, error)
	DeleteSong(string, string) (bool, error)
	UpdateSong(EffectiveMobile.Song) (bool, error)
	Add(EffectiveMobile.Song) (int, error)
}
type Service struct {
	Song
}
