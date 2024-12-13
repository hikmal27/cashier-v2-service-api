package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
)

type MenuService struct {
	repository repository.MenuRepository
}

func NewMenuService(r repository.MenuRepository) MenuService {
	return MenuService{
		repository: r,
	}
}

// @Summary : List menu
// @Description : Get menus from repository
// @Author : azrielfatur
func (s *MenuService) List(param bool) ([]entity.Menu, error) {
	menus, err := s.repository.FindAll(param)

	if err != nil {
		return menus, err
	}

	return menus, nil
}

// @Summary : Insert role
// @Description : Get role to repository
// @Author : azrielfatur
func (s *MenuService) Insert(req request.MenuRequest) (entity.Menu, error) {
	menu := entity.Menu{
		Name:         req.Name,
		Url:          req.Url,
		Icon:         req.Icon,
		Ord:          req.Ord,
		IsHeader:     req.IsHeader,
		ParentID:     req.ParentID,
		CreateUserID: req.CreatedBy,
	}

	newMenu, err := s.repository.Insert(menu)

	if err != nil {
		return newMenu, err
	}

	return newMenu, nil
}

// @Summary : Find role
// @Description : Find role by id from repository
// @Author : azrielfatur
func (s *MenuService) FindById(ID int) (entity.Menu, error) {
	menu, err := s.repository.FindById(ID)

	if err != nil {
		return menu, err
	}

	return menu, nil
}

// @Summary : Update role
// @Description : Update role by id from repository
// @Author : azrielfatur
func (s *MenuService) Update(req request.MenuRequest, ID int) (entity.Menu, error) {
	menu := entity.Menu{
		Name:         req.Name,
		Url:          req.Url,
		Icon:         req.Icon,
		Ord:          req.Ord,
		ParentID:     req.ParentID,
		IsHeader:     req.IsHeader,
		UpdateUserID: req.UpdatedBy,
	}

	updateMenu, err := s.repository.Update(menu, ID)

	if err != nil {
		return updateMenu, err
	}

	return updateMenu, nil
}

// @Summary : Delete role
// @Description : Delete role to repository
// @Author : azrielfatur
func (s *MenuService) Delete(ID int) (bool, error) {
	menu, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return menu, nil
}
