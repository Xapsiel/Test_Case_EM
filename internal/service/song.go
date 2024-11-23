package service

import (
	"github.com/Xapsiel/EffectiveMobile/internal/models"
	"github.com/Xapsiel/EffectiveMobile/internal/repository"
	"strings"
)

type SongService struct {
	repo repository.Song
}

func NewSongService(repo repository.Song) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) GetSongs(filter models.Filter) ([]models.Song, error) {
	return s.repo.GetSongs(filter)
}
func (s *SongService) GetSongVerse(song models.Song, VerseNumber int) (string, int, error) {
	text, id, err := s.repo.GetSongVerse(song)
	if err != nil {
		return "", 0, err
	}
	VerseNumber -= 1
	textParts := strings.Split(text, "\n\n")
	if len(textParts)-1 < VerseNumber || VerseNumber < 0 {
		return textParts[len(textParts)-1], id, nil // возвращаем последний элемент
	}
	return textParts[VerseNumber], id, nil
}
func (s *SongService) DeleteSong(song models.Song) (bool, error) {
	return s.repo.DeleteSong(song)
}
func (s *SongService) UpdateSong(song models.Song) (bool, models.Song, error) {
	return s.repo.UpdateSong(song)
}
func (s *SongService) Add(song models.Song) (bool, int, error) {
	return s.repo.Add(song)
}

/*
	GetSongs(Filter) ([]EffectiveMobile.Song, error)
	GetSongVerse(string, string, int) (string, error)
	DeleteSong(string, string) (bool, error)
	UpdateSong(EffectiveMobile.Song) (EffectiveMobile.Song, error)
	Add(song EffectiveMobile.Song) (int, error)
*/
