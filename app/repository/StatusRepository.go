package repository

import (
	"gorbac/app/entity"
	"gorbac/config"
)

type StatusRepository struct {
	config config.Database
}

func NewStatusRepository(database config.Database) StatusRepository {
	return StatusRepository{
		config: database,
	}
}

// @Summary : Get status
// @Description : -
// @Author : rasmadibnu
func (r *StatusRepository) FindAll() ([]entity.Status, error) {
	var status []entity.Status

	err := r.config.DB.Find(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

// @Summary : Insert status
// @Description : Insert status to database
// @Author : rasmadibnu
func (r *StatusRepository) Insert(status entity.Status) (entity.Status, error) {
	err := r.config.DB.Create(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

// @Summary : Get status
// @Description : find status by ID
// @Author : rasmadibnu
func (r *StatusRepository) FindById(ID int) (entity.Status, error) {
	var status entity.Status

	err := r.config.DB.Where("id = ?", ID).First(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

// @Summary : Update status
// @Description : Update status by ID
// @Author : rasmadibnu
func (r *StatusRepository) Update(status entity.Status, ID int) (entity.Status, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

// @Summary : Delete status
// @Description : Delete status temporary
// @Author : rasmadibnu
func (r *StatusRepository) Delete(ID int) (bool, error) {
	var status entity.Status

	err := r.config.DB.Where("id = ?", ID).Delete(&status).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
