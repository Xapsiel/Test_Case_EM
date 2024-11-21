package EffectiveMobile

import "time"

type Song struct {
	ReleaseDate time.Time `json:"release_date" `
	Text        *string   `json:"release_date,omitempty" `
	Group       string    `json:"release_date,omitempty" `
	SongName    string    `json:"release_date,omitempty" `
	Link        *string   `json:"release_date,omitempty" `
}
