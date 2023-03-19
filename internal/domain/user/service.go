package user

import (
	"context"
	"errors"
	"jwt_registration_api/internal/adapters/api/http_handlers/dto"
	"jwt_registration_api/internal/adapters/api/http_handlers/user"
	"jwt_registration_api/internal/domain/regJwt"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	storage Storage
	jwtServ *regJwt.JwtServ
}

func NewService(storage Storage, jwtServ *regJwt.JwtServ) user.Service {
	return &service{
		storage: storage,
		jwtServ: jwtServ,
	}
}

func (s *service) Register(ctx context.Context, regInput *dto.RegisterInput) (*dto.RegisterPayload, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(regInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Can`t hash password: " + err.Error())
	}

	user := &User{
		Login:       regInput.Login,
		Email:       regInput.Email,
		Password:    string(hashPass),
		PhoneNumber: regInput.PhoneNumber,
	}
	// Считаем, что данные от пользователя прошли валидацию на фронте
	user, err = s.storage.InsertUser(user)
	if err != nil {
		return nil, errors.New("Can`t insert user: " + err.Error())
	}

	token, err := s.jwtServ.GenerateUserToken(user.Id)
	if err != nil {
		return nil, errors.New("Can`t generate user token: " + err.Error())
	}

	regPayload := &dto.RegisterPayload{
		Id:          user.Id,
		Login:       user.Login,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Token:       token,
	}

	return regPayload, nil
}

func (s *service) Login(ctx context.Context, loginInput *dto.LoginInput) (*dto.LoginPayload, error) {
	user, err := s.storage.GetUserByLogin(loginInput.Login)
	if err != nil {
		return nil, errors.New("The user with this username was not found: " + err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		return nil, errors.New("The password is not correct: " + err.Error())
	}

	token, err := s.jwtServ.GenerateUserToken(user.Id)
	if err != nil {
		return nil, errors.New("Can`t generate user token: " + err.Error())
	}

	loginPayload := &dto.LoginPayload{
		Token: token,
	}

	return loginPayload, nil
}
