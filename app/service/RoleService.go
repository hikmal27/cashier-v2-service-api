package service

import (
	"fmt"
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
)

type RoleService struct {
	repository repository.RoleRepository
}

func NewRoleService(r repository.RoleRepository) RoleService {
	return RoleService{
		repository: r,
	}
}

// @Summary : List role
// @Description : Get roles from repository
// @Author : azrielfatur
func (s *RoleService) List() ([]entity.Role, error) {
	roles, err := s.repository.FindAll()

	if err != nil {
		return roles, err
	}

	return roles, nil
}

// @Summary : Insert role
// @Description : insert role to repository
// @Author : azrielfatur
func (s *RoleService) Insert(req request.RoleRequest) (entity.Role, error) {
	role := entity.Role{
		Name:         req.Name,
		Description:  req.Description,
		CreateUserID: req.CreatedBy,
		UpdateUserID: req.UpdatedBy,
	}

	newRole, err := s.repository.Insert(role)

	if err != nil {
		return newRole, err
	}

	return newRole, nil
}

// @Summary : Find role
// @Description : Find role by id from repository
// @Author : azrielfatur
func (s *RoleService) FindById(ID int) (entity.Role, error) {
	role, err := s.repository.FindById(ID)

	if err != nil {
		return role, err
	}

	return role, nil
}

// @Summary : Update role
// @Description : Update role by id from repository
// @Author : azrielfatur
func (s *RoleService) Update(req request.RoleRequest, ID int) (entity.Role, error) {
	role := entity.Role{
		Name:         req.Name,
		Description:  req.Description,
		CreateUserID: req.CreatedBy,
		UpdateUserID: req.UpdatedBy,
	}

	fmt.Print(role)

	updateRole, err := s.repository.Update(role, ID)

	if err != nil {
		return updateRole, err
	}

	return updateRole, nil
}

// @Summary : Delete role
// @Description : Delete role to repository
// @Author : azrielfatur
func (s *RoleService) Delete(ID int) (bool, error) {
	role, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return role, nil
}
