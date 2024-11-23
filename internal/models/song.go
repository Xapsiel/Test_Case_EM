package models

type Song struct {
	ID          int    `json:"id,omitempty" example:"1" db:"id"`
	Group       string `json:"group,omitempty" example:"Muse" db:"group_name"`
	SongName    string `json:"song_name,omitempty" example:"Supermassive Black Hole" db:"song_name"`
	ReleaseDate string `json:"release_date,omitempty" example:"2006-07-19" db:"release_date"`
	Link        string `json:"link,omitempty" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw" db:"link"`
	Text        string `json:"text,omitempty" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?" db:"text"`
}
