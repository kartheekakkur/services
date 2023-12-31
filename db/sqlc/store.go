package db

import (
	"database/sql"
)

type Store struct {
	*Queries
	db *sql.DB
}

//NewStore creates new store

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
