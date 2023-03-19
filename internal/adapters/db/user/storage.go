package user

import (
	"database/sql"
	"errors"
	"jwt_registration_api/internal/domain/user"

	qb "github.com/Masterminds/squirrel"
)

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) user.Storage {
	return &storage{db: db}
}

func (s *storage) InsertUser(newUser *user.User) (*user.User, error) {
	query := qb.Insert(TableAuthUser).
		Columns(ColumnsAuthUser).
		Values(newUser.Login, newUser.Email, newUser.Password, newUser.PhoneNumber).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(qb.Dollar).
		RunWith(s.db)

	err := query.QueryRow().Scan(&newUser.Id)
	if err != nil {
		return nil, errors.New("Can`t insert user into auth_user: " + err.Error())
	}
	return newUser, nil
}

func (s *storage) GetUserByLogin(login string) (*user.User, error) {
	query := qb.
		Select(SelectAllAuthUser).
		From(TableAuthUser).
		Where(qb.Eq{"login": login}).
		PlaceholderFormat(qb.Dollar).
		RunWith(s.db)

	var user user.User
	err := query.QueryRow().Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.PhoneNumber)
	if err != nil {
		return nil, errors.New("Can`t find user with this login: " + err.Error())
	}

	return &user, nil
}
