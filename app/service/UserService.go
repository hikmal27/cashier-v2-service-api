package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
	"gorbac/security"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

// @Summary : List User
// @Description : Get users from repository
// @Author : azrielfatur
func (s *UserService) List() ([]entity.User, error) {
	users, err := s.repository.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}

// @Summary : Insert user
// @Description : Insert user to repository
// @Author : azrielfatur
func (s *UserService) Insert(req request.RegisterRequest) (entity.User, error) {
	hash, err := security.HashPassword(req.Password)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		// GroupID:      req.GroupID,
		Username:     req.Username,
		Name:         req.Name,
		Email:        req.Email,
		Password:     hash,
		CreateUserID: req.CreatedBy,
	}

	newUser, err := s.repository.Insert(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// @Summary : Find user
// @Description : Find user by username
// @Author : azrielfatur
func (s *UserService) FindByUsername(username string) (entity.User, error) {
	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Find user
// @Description : Find user by id
// @Author : azrielfatur
func (s *UserService) FindById(ID int) (entity.User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Udate user
// @Description : Update user to repository
// @Author : azrielfatur
func (s *UserService) Update(req request.UserUpdateRequest, ID int) (entity.User, error) {
	hash, err := security.HashPassword(req.Password)

	if err != nil {
		return entity.User{}, err
	}

	if req.Password != "" {
		user := entity.User{
			// GroupID:      req.GroupID,
			Username:     req.Username,
			Name:         req.Name,
			Email:        req.Email,
			Password:     hash,
			UpdateUserID: req.UpdatedBy,
		}

		updateUser, err := s.repository.Update(user, ID)
		if err != nil {
			return updateUser, err
		}

		return updateUser, nil
	} else {
		user := entity.User{
			// GroupID:      req.GroupID,
			Username:     req.Username,
			Name:         req.Name,
			Email:        req.Email,
			UpdateUserID: req.UpdatedBy,
		}

		updateUser, err := s.repository.Update(user, ID)
		if err != nil {
			return updateUser, err
		}

		return updateUser, nil
	}
}

// @Summary : Delete user
// @Description : Delete user to repository
// @Author : azrielfatur
func (s *UserService) Delete(ID int) (bool, error) {
	user, err := s.repository.Delete(ID)

	if err != nil {
		return false, nil
	}

	return user, nil
}
