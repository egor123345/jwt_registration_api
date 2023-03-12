package user

type Storage interface {
	InsertUser(user *User) (*User, error)
	GetUserByLogin(login string) (*User, error)
}
