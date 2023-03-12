package user

import (
	"database/sql"
	"jwt_registration_api/internal/domain/user"
)

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) user.Storage {
	return &storage{db: db}
}

func (s *storage) InsertUser(user *user.User) (*user.User, error) {
	return nil, nil
}

func (s *storage) GetUserByLogin(login string) (*user.User, error) {
	return nil, nil
}
