package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type AreaRepository struct {
	config config.Database
}

func NewAreaRepository(database config.Database) AreaRepository {
	return AreaRepository{
		config: database,
	}
}

// @Summary : Get area
// @Description : find all area
// @Author : rasmadibnu
func (r *AreaRepository) FindAll() ([]entity.BranchArea, error) {
	var area []entity.BranchArea

	err := r.config.DB.Preload("Branch").Where("areaID not in ('9999','999')").Or("areaID is null").Find(&area).Error

	if err != nil {
		return area, err
	}

	return area, nil
}

// @Summary : Get area
// @Description : find area by ID
// @Author : rasmadibnu
func (r *AreaRepository) FindById(ID int) (entity.BranchArea, error) {
	var area entity.BranchArea

	err := r.config.DB.Preload("Branch").Where("areaID not in ('9999','999')").Or("areaID is null").Where("areaID = ?", ID).Find(&area).Error

	if err != nil {
		return area, err
	}

	return area, nil
}
