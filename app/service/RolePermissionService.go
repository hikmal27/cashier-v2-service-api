package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
)

type RolePermissionService struct {
	repository repository.RolePermissionRepository
}

func NewRolePermissionService(r repository.RolePermissionRepository) RolePermissionService {
	return RolePermissionService{
		repository: r,
	}
}

// @Summary : Attach menu
// @Description : Attach user has menu
// @Author : azrielfatur
func (s *RolePermissionService) AttachPermission(req request.RoleHasPermission) ([]entity.RolePermission, error) {
	attach, err := s.repository.Attach(req.Permission)

	if err != nil {
		return attach, err
	}

	return attach, nil
}

// @Summary : Detach menu
// @Description : Detach role has menus
// @Author : azrielfatur
func (s *RolePermissionService) DetachPermission(req request.RoleHasPermission) (bool, error) {
	attach, err := s.repository.Detach(req.Permission)

	if err != nil {
		return attach, err
	}

	return attach, nil
}
