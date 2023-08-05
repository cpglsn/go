package data

import "time"

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // with json:"-" we avoid displaying this data in the json
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"` // with ,omitempty we make this field to be
	Pages     int       `json:"pages,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Version   int32     `json:"-"`
}
