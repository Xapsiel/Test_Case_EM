package handler

type Song struct {
	ID          int    `json:"id" db:"id"`
	ReleaseDate string `json:"release_date" db:"release_date" `
	Text        string `json:"text,omitempty" db:"text"`
	Group       string `json:"group,omitempty" db:"group_name"`
	SongName    string `json:"song_name,omitempty" db:"song_name"`
	Link        string `json:"link,omitempty" db:"link"`
}
