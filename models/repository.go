package models

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) Find(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM movies WHERE id = $1`
	rows := repo.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate,
		&movie.Runtime, &movie.Runtime, &movie.MPAARating, &movie.CreatedAt, &movie.UpdateAt)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repo *Repository) FindAll() ([]*Movie, error) {
	return nil, nil
}
