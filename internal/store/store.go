package store

import (
	"context"
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("Record Not Found")

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetPostById(context.Context, int64) (*Post, error)
	}
	User interface {
		Create(context.Context, *User) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		User:  &UsersStore{db},
	}
}
