package models

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) FindMovie(id int) (*Movie, error) {
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
		return nil, err
	}
	defer rows.Close()

	mgs := make(map[int]string)

	for rows.Next() {
		var mg MovieGenre
		err = rows.Scan(&mg.ID, &mg.MovieID, &mg.GenreID, &mg.CreatedAt, &mg.UpdateAt, &mg.Genre.GenreName)
		if err != nil {
			return nil, err
		}
		mgs[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = mgs

	return &movie, nil
}

func (repo *Repository) FindMovieByGenre(genre string) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT m.id, title, m.description, m.year, m.release_date, m.runtime, m.rating, m.mpaa_rating 
		FROM movies m INNER JOIN movies_genres mg ON m.id = mg.movie_id INNER JOIN genres g 
        ON g.id = mg.genre_id WHERE g.genre_name = $1`

	rows, err := repo.DB.QueryContext(ctx, query, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var m Movie
		err = rows.Scan(&m.ID, &m.Title, &m.Description, &m.Year, &m.ReleaseDate, &m.Runtime, &m.Rating, &m.MPAARating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &m)
	}

	return movies, nil
}

func (repo *Repository) FindMovieAll() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating
			FROM movies ORDER BY title`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Runtime, &movie.Rating, &movie.MPAARating)
		if err != nil {
			return nil, err
		}

		query = `SELECT mg.id, mg.movie_id, mg.genre_id, mg.created_at, mg.updated_at, g.genre_name 
		FROM movies_genres mg LEFT JOIN genres g ON g.id = mg.genre_id WHERE mg.movie_id = $1`
		genreRows, err := repo.DB.QueryContext(ctx, query, movie.ID)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()

		mgs := make(map[int]string)

		for genreRows.Next() {
			var mg MovieGenre
			err = genreRows.Scan(&mg.ID, &mg.MovieID, &mg.GenreID, &mg.CreatedAt, &mg.UpdateAt, &mg.Genre.GenreName)
			if err != nil {
				return nil, err
			}
			mgs[mg.ID] = mg.Genre.GenreName
		}

		movie.MovieGenre = mgs

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (repo *Repository) FindGenresAll() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, genre_name
			FROM genres ORDER BY genre_name`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var g Genre
		err = rows.Scan(&g.ID, &g.GenreName)
		if err != nil {
			return nil, err
		}

		genres = append(genres, &g)
	}

	return genres, nil
}
