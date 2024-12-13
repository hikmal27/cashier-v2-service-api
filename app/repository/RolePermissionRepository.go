package repository

import (
	"gorbac/app/entity"
	"gorbac/config"

	"gorm.io/gorm/clause"
)

type RolePermissionRepository struct {
	config config.Database
}

func NewRolePermissionRepository(database config.Database) RolePermissionRepository {
	return RolePermissionRepository{
		config: database,
	}
}

// @Summary : Assign role has menu
// @Description : Assign role has menu to database
// @Author : azrielfatur
func (r *RolePermissionRepository) Attach(rolePermission []entity.RolePermission) ([]entity.RolePermission, error) {
	err := r.config.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"create", "read", "update", "delete", "updated_by"}),
	}).Create(&rolePermission).Error

	if err != nil {
		return rolePermission, err
	}

	return rolePermission, nil
}

// @Summary : Un-Assign role has menu
// @Description : Un-Assign role has menu to database
// @Author : azrielfatur
func (r *RolePermissionRepository) Detach(rolePermission []entity.RolePermission) (bool, error) {
	err := r.config.DB.Unscoped().Delete(&rolePermission).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
