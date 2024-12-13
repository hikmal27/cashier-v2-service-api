package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type RoleMenuRepository struct {
	config config.Database
}

func NewRoleMenuRepository(database config.Database) RoleMenuRepository {
	return RoleMenuRepository{
		config: database,
	}
}

// @Summary : Find role has menu
// @Description : Find role has menu to database
// @Author : azrielfatur
func (r *RoleMenuRepository) Find(RoleID int) (entity.UserRole, error) {
	var userRole entity.UserRole

	err := r.config.DB.Where(entity.UserRole{RoleID: RoleID}).First(&userRole).Error

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary : Assign role has menu
// @Description : Assign role has menu to database
// @Author : azrielfatur
func (r *RoleMenuRepository) Attach(roleMenu []entity.RoleMenu) ([]entity.RoleMenu, error) {
	err := r.config.DB.Create(&roleMenu).Error

	if err != nil {
		return roleMenu, err
	}

	return roleMenu, nil
}

// @Summary : Un-Assign role has menu
// @Description : Un-Assign role has menu to database
// @Author : azrielfatur
func (r *RoleMenuRepository) Detach(roleMenu []entity.RoleMenu) ([]entity.RoleMenu, error) {
	var req []entity.RoleMenu

	RoleID := make([]int, 0)
	MenuID := make([]int, 0)

	for _, col := range roleMenu {
		RoleID = append(RoleID, col.RoleID)
		MenuID = append(MenuID, col.MenuID)
	}

	err := r.config.DB.Where("role_id IN ?", RoleID).Delete(&roleMenu).Error
	if err != nil {
		return req, err
	}

	return req, nil
}
