package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
)

type UserRoleService struct {
	repository repository.UserRoleRepository
}

func NewUserRoleService(r repository.UserRoleRepository) UserRoleService {
	return UserRoleService{
		repository: r,
	}
}

// @Summary : Find role
// @Description : Find user has role
// @Author : azrielfatur
func (s *UserRoleService) FindRole(req request.UserRole) (entity.UserRole, error) {
	userRole := entity.UserRole{
		UserID: req.UserID,
		RoleID: req.RoleID,
	}
	userRole, err := s.repository.Find(userRole)

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary : Attach role
// @Description : Attach user has role
// @Author : azrielfatur
func (s *UserRoleService) AttachRole(req request.UserHasRole) ([]entity.UserRole, error) {
	dettach, err := s.repository.Detach(req.Roles)
	attach, err1 := s.repository.Attach(req.Roles)

	if err != nil {
		return dettach, err
	}

	if err1 != nil {
		return attach, err
	}

	return attach, nil
}

// @Summary : Detach role
// @Description : Detach user has role
// @Author : azrielfatur
// func (s *UserRoleService) DetachRole(req request.UserHasRole) (bool, error) {
// 	detach, err := s.repository.Detach(req.Roles)

// 	if err != nil {
// 		return false, err
// 	}

// 	return detach, nil
// }
