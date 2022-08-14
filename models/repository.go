package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

func (repo *Repository) FindMovieAll(genreId ...int) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genreId) > 0 {
		where = fmt.Sprintf("WHERE id IN (SELECT movie_id FROM movies_genres WHERE genre_id = %v)", genreId[0])
	}

	query := fmt.Sprintf(`SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating
			FROM movies %s ORDER BY title`, where)
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

func (repo *Repository) InsertMovie(m Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO movies(title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = repo.DB.ExecContext(ctx, query, m.Title, m.Description, m.Year, m.ReleaseDate, m.Runtime, m.Rating, m.MPAARating, m.CreatedAt, m.UpdateAt)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
