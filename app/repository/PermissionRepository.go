package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type PermissionRepository struct {
	config config.Database
}

func NewPermissionRepository(database config.Database) PermissionRepository {
	return PermissionRepository{
		config: database,
	}
}

// @Summary : Insert permission
// @Description : Insert permission to database
// @Author : azrielfatur
func (r *PermissionRepository) Insert(permission entity.Permission) (entity.Permission, error) {
	err := r.config.DB.Debug().Create(&permission).Error

	if err != nil {
		return permission, err
	}

	return permission, nil
}

// @Summary : Get permissions
// @Description : -
// @Author : azrielfatur
func (r *PermissionRepository) FindAll() ([]entity.Permission, error) {
	var permissions []entity.Permission

	err := r.config.DB.Debug().Find(&permissions).Error

	if err != nil {
		return permissions, err
	}

	return permissions, nil
}

// @Summary : Get permission
// @Description : Find permission by ID
// @Author : azrielfatur
func (r *PermissionRepository) FindById(ID int) (entity.Permission, error) {
	var permission entity.Permission

	err := r.config.DB.Debug().Where("id = ?", ID).First(&permission).Error

	if err != nil {
		return permission, err
	}

	return permission, nil
}

// @Summary : Update permission
// @Description : Update permission by ID
// @Author : azrielfatur
func (r *PermissionRepository) Update(permission entity.Permission, ID int) (entity.Permission, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&permission).Error

	if err != nil {
		return permission, err
	}

	return permission, nil
}

// @Summary : Delete permission
// @Description : Delete permission temporary
// @Author : azrielfatur
func (r *PermissionRepository) Delete(ID int) (bool, error) {
	var permission entity.Permission

	err := r.config.DB.Debug().Where("id = ?", ID).Delete(&permission).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
