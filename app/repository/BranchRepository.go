package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type BranchRepository struct {
	config config.Database
}

func NewBranchRepository(database config.Database) BranchRepository {
	return BranchRepository{
		config: database,
	}
}

// @Summary : Get branch
// @Description : find all branch
// @Author : rasmadibnu
func (r *BranchRepository) FindAll(param map[string]interface{}) ([]entity.Branch, error) {
	var branch []entity.Branch

	err := r.config.DB.Where(param).Preload("BranchArea").Preload("Warehouses").Not(map[string]interface{}{"areaID": []int{9999, 999, 0}}).Find(&branch).Error

	if err != nil {
		return branch, err
	}

	return branch, nil
}

// @Summary : Get branch
// @Description : find branch by ID
// @Author : rasmadibnu
func (r *BranchRepository) FindById(ID int) (entity.Branch, error) {
	var branch entity.Branch

	err := r.config.DB.Preload("BranchArea").Preload("Warehouses").Not(map[string]interface{}{"areaID": []int{9999, 999, 0}}).Where("branchID = ?", ID).First(&branch).Error

	if err != nil {
		return branch, err
	}

	return branch, nil
}
