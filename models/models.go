package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB Repository
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: Repository{db},
	}
}

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"releaseDate"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdateAt    time.Time    `json:"update_at"`
	MovieGenre  []MovieGenre `json:"-"`
}

type Genre struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	GenreID   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
