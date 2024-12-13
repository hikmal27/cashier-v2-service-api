package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
)

type PermissionService struct {
	repository repository.PermissionRepository
}

func NewPermissionService(r repository.PermissionRepository) PermissionService {
	return PermissionService{
		repository: r,
	}
}

// @Summary : List permission
// @Description : Get permissions from repository
// @Author : azrielfatur
func (s *PermissionService) List() ([]entity.Permission, error) {
	permissions, err := s.repository.FindAll()

	if err != nil {
		return permissions, err
	}

	return permissions, nil
}

// @Summary : Insert permission
// @Description : insert permission to repository
// @Author : azrielfatur
func (s *PermissionService) Insert(req request.PermissionRequest) (entity.Permission, error) {
	permission := entity.Permission{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}

	newPermission, err := s.repository.Insert(permission)

	if err != nil {
		return newPermission, err
	}

	return newPermission, nil
}

// @Summary : Find permission
// @Description : Find permission by id from repository
// @Author : azrielfatur
func (s *PermissionService) FindById(ID int) (entity.Permission, error) {
	permission, err := s.repository.FindById(ID)

	if err != nil {
		return permission, err
	}

	return permission, nil
}

// @Summary : Update permission
// @Description : Update permission by id from repository
// @Author : azrielfatur
func (s *PermissionService) Update(req request.PermissionRequest, ID int) (entity.Permission, error) {
	permission := entity.Permission{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}

	updatePermission, err := s.repository.Update(permission, ID)

	if err != nil {
		return updatePermission, err
	}

	return updatePermission, nil
}

// @Summary : Delete permission
// @Description : Delete permission to repository
// @Author : azrielfatur
func (s *PermissionService) Delete(ID int) (bool, error) {
	permission, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return permission, nil
}
