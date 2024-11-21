package service

import (
	"github.com/Xapsiel/EffectiveMobile"
	"github.com/Xapsiel/EffectiveMobile/internal/repository"
	"strings"
)

type SongService struct {
	repo repository.Song
}

func NewSongService(repo repository.Song) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) GetSongs(filter EffectiveMobile.Filter) ([]EffectiveMobile.Song, error) {
	return s.repo.GetSongs(filter)
}
func (s *SongService) GetSongVerse(song string, group string, VerseNumber int) (string, error) {
	text, err := s.repo.GetSongVerse(song, group)
	if err != nil {
		return "", err
	}
	textParts := strings.Split(text, "\n\n")
	if len(textParts)-1 < VerseNumber {
		return textParts[len(textParts)-1], nil // возвращаем последний элемент
	}
	return textParts[VerseNumber], nil
}
func (s *SongService) DeleteSong(song, group string) (bool, error) {
	return s.repo.DeleteSong(song, group)
}
func (s *SongService) UpdateSong(song EffectiveMobile.Song) (bool, error) {
	return s.repo.UpdateSong(song)
}
func (s *SongService) Add(song EffectiveMobile.Song) (int, error) {
	return s.repo.Add(song)
}

/*
	GetSongs(Filter) ([]EffectiveMobile.Song, error)
	GetSongVerse(string, string, int) (string, error)
	DeleteSong(string, string) (bool, error)
	UpdateSong(EffectiveMobile.Song) (EffectiveMobile.Song, error)
	Add(song EffectiveMobile.Song) (int, error)
*/
