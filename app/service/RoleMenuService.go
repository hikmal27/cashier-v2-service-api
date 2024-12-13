package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
)

type RoleMenuService struct {
	repository repository.RoleMenuRepository
}

func NewRoleMenuService(r repository.RoleMenuRepository) RoleMenuService {
	return RoleMenuService{
		repository: r,
	}
}

// @Summary : Find role
// @Description : Find user has role
// @Author : azrielfatur
func (s *RoleMenuService) FindMenu(RoleID int) (entity.UserRole, error) {
	userRole, err := s.repository.Find(RoleID)

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary : Attach menu
// @Description : Attach user has menu
// @Author : azrielfatur
func (s *RoleMenuService) AttachMenu(req request.RoleHasMenu) ([]entity.RoleMenu, error) {
	dettach, err := s.repository.Detach(req.Menus)
	attach, err1 := s.repository.Attach(req.Menus)

	if err != nil {
		return dettach, err
	}

	if err1 != nil {
		return attach, err1
	}

	return attach, nil
}

// @Summary : Detach menu
// @Description : Detach role has menus
// @Author : azrielfatur
func (s *RoleMenuService) DetachMenu(req request.RoleHasMenu) ([]entity.RoleMenu, error) {
	attach, err := s.repository.Detach(req.Menus)

	if err != nil {
		return attach, err
	}

	return attach, nil
}
