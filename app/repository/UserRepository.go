package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type UserRepository struct {
	config config.Database
}

func NewUserRepository(database config.Database) UserRepository {
	return UserRepository{
		config: database,
	}
}

// @Summary : Insert user
// @Description : Insert user to database
// @Author : azrielfatur
func (r *UserRepository) Insert(user entity.User) (entity.User, error) {
	err := r.config.DB.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Get users
// @Description : -
// @Author : azrielfatur
func (r *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := r.config.DB.Preload("Roles.Menus").Preload("CreatedBy").Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

// @Summary : Get user
// @Description : Find user by Username
// @Author : azrielfatur
func (r *UserRepository) FindByUsername(username string) (entity.User, error) {
	var user entity.User

	r.config.DB.Preload("Roles.Permissions").Preload("Roles.Menus.Parent").Where("username = ?", username).First(&user)

	return user, nil
}

// @Summary : Get user
// @Description : Find user by ID
// @Author : azrielfatur
func (r *UserRepository) FindById(ID int) (entity.User, error) {
	var user entity.User
	err := r.config.DB.Preload("Roles.Permissions").Preload("Roles.Menus.Parent").Preload("CreatedBy").Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Update user
// @Description : Update user
// @Author : azrielfatur
func (r *UserRepository) Update(user entity.User, ID int) (entity.User, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Delete user
// @Description : Delete user temporary
// @Author : azrielfatur
func (r *UserRepository) Delete(ID int) (bool, error) {
	var user entity.User

	err := r.config.DB.Where("id = ?", ID).Delete(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// @Summary : Delete user
// @Description : Delete user permanent
// @Author : azrielfatur
func (r *UserRepository) ForceDelete(ID int) (bool, error) {
	var user entity.User

	err := r.config.DB.Unscoped().Where("id = ?", ID).Delete(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
