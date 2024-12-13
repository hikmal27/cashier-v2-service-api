package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
	"gorbac/security"
)

type AuthService struct {
	repository repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return AuthService{
		repository: r,
	}
}

func (s *AuthService) Login(req request.LoginRequest) (entity.User, error) {
	reqUsername := req.Username
	reqPassword := req.Password

	user, err := s.repository.FindByUsername(reqUsername)

	if err != nil {
		return user, err
	}

	checkPassword := security.CheckPassword(user.Password, reqPassword)

	if checkPassword == nil {
		return user, nil
	}

	return user, checkPassword
}
