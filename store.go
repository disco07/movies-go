package main

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Store interface {
	Open() error
	Close() error
}

type dbStore struct {
	db *sqlx.DB
}

func (s dbStore) Open() error {
	db, err := sqlx.Connect("mysql", "user=root dbname=formationplus sslmode=disable")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	s.db = db
	return nil
}

func (s dbStore) Close() error {
	return s.db.Close()
}
