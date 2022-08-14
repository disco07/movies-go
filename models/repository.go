package models

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) Find(id int) (*Movie, error) {
	return nil, nil
}

func (repo *Repository) FindAll() ([]*Movie, error) {
	return nil, nil
}
