package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type UserRoleRepository struct {
	config config.Database
}

func NewUserRoleRepository(database config.Database) UserRoleRepository {
	return UserRoleRepository{
		config: database,
	}
}

// @Summary : Find user has role
// @Description : Find user has role to database
// @Author : azrielfatur
func (r *UserRoleRepository) Find(userRole entity.UserRole) (entity.UserRole, error) {
	err := r.config.DB.Where("user_id = ? AND role_id = ?", userRole.UserID, userRole.RoleID).First(&userRole).Error

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary : Assign user has role
// @Description : Assign user has role to database
// @Author : azrielfatur
func (r *UserRoleRepository) Attach(userRole []entity.UserRole) ([]entity.UserRole, error) {
	err := r.config.DB.Create(&userRole).Error

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary : Un-Assign user has role
// @Description : Un-Assign user has role to database
// @Author : azrielfatur
func (r *UserRoleRepository) Detach(userRole []entity.UserRole) ([]entity.UserRole, error) {
	var req []entity.UserRole
	UserID := make([]int, 0)
	RoleID := make([]int, 0)

	for _, col := range userRole {
		UserID = append(UserID, col.UserID)
		RoleID = append(RoleID, col.RoleID)
	}
	err := r.config.DB.Unscoped().Where("user_id = ?", userRole[0].UserID).Delete(&userRole).Error

	if err != nil {
		return req, err
	}

	return req, nil
}
