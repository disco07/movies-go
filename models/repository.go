package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) Find(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM movies WHERE id = $1`
	row := repo.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate,
		&movie.Runtime, &movie.Runtime, &movie.MPAARating, &movie.CreatedAt, &movie.UpdateAt)
	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.movie_id, mg.genre_id, mg.created_at, mg.updated_at, g.genre_name 
		FROM movies_genres mg LEFT JOIN genres g ON g.id = mg.genre_id WHERE mg.movie_id = $1`
	rows, err := repo.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mgs []MovieGenre

	for rows.Next() {
		var mg MovieGenre
		err = rows.Scan(&mg.ID, &mg.MovieID, &mg.GenreID, &mg.CreatedAt, &mg.UpdateAt, &mg.Genre.GenreName)
		if err != nil {
			log.Fatal(err)
		}
		mgs = append(mgs, mg)
	}

	movie.MovieGenre = mgs

	return &movie, nil
}

func (repo *Repository) FindAll() ([]*Movie, error) {
	return nil, nil
}
