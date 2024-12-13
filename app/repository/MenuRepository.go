package repository

import (
	"gorbac/app/entity"
	"gorbac/config"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository struct {
	config config.Database
}

func NewMenuRepository(database config.Database) MenuRepository {
	return MenuRepository{
		config: database,
	}
}

// @Summary : Insert menu
// @Description : Insert menu to database
// @Author : azrielfatur
func (r *MenuRepository) Insert(menu entity.Menu) (entity.Menu, error) {
	err := r.config.DB.Create(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

// @Summary : Get menus
// @Description : -
// @Author : azrielfatur
func (r *MenuRepository) FindAll(param bool) ([]entity.Menu, error) {
	var menus []entity.Menu

	if param {
		err := r.config.DB.Preload("Roles").Order("ord ASC").Preload("Childs", func(db *gorm.DB) *gorm.DB {
			return db.Order("ord ASC").Preload("Parent").Preload("Roles")
		}).Preload(clause.Associations).Preload("Childs", func(db *gorm.DB) *gorm.DB {
			return db.Order("ord ASC").Preload("Roles")
		}).Preload(clause.Associations).Preload("Childs", func(db *gorm.DB) *gorm.DB {
			return db.Order("ord ASC").Preload("Roles")
		}).Preload(clause.Associations).Preload("Childs.Childs", func(db *gorm.DB) *gorm.DB {
			return db.Order("ord ASC").Preload("Roles")
		}).Find(&menus).Error

		if err != nil {
			return menus, err
		}

		return menus, nil
	}

	err := r.config.DB.Preload("Roles").Order("ord ASC").Preload("Childs", func(db *gorm.DB) *gorm.DB {
		return db.Order("ord ASC").Preload("Parent").Preload("Roles")
	}).Preload(clause.Associations).Preload("Childs", func(db *gorm.DB) *gorm.DB {
		return db.Order("ord ASC").Preload("Roles")
	}).Preload(clause.Associations).Preload("Childs", func(db *gorm.DB) *gorm.DB {
		return db.Order("ord ASC").Preload("Roles")
	}).Preload(clause.Associations).Preload("Childs.Childs", func(db *gorm.DB) *gorm.DB {
		return db.Order("ord ASC").Preload("Roles")
	}).Where("parent_id", nil).Find(&menus).Error

	if err != nil {
		return menus, err
	}

	return menus, nil
}

// @Summary : Get menu
// @Description : find menu by ID
// @Author : azrielfatur
func (r *MenuRepository) FindById(ID int) (entity.Menu, error) {
	var menu entity.Menu

	err := r.config.DB.Where("id = ?", ID).First(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

// @Summary : Update menu
// @Description : Update menu by ID
// @Author : azrielfatur
func (r *MenuRepository) Update(menu entity.Menu, ID int) (entity.Menu, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

// @Summary : Delete menu
// @Description : Delete menu temporary
// @Author : azrielfatur
func (r *MenuRepository) Delete(ID int) (bool, error) {
	var menu entity.Menu

	err := r.config.DB.Where("id = ?", ID).Delete(&menu).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
