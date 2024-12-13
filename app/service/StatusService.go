package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
	"gorbac/app/request"
)

type StatusService struct {
	repository repository.StatusRepository
}

func NewStatusService(r repository.StatusRepository) StatusService {
	return StatusService{
		repository: r,
	}
}

// @Summary : List status
// @Description : Get status from repository
// @Author : rasmadibnu
func (s *StatusService) List() ([]entity.Status, error) {
	status, err := s.repository.FindAll()

	if err != nil {
		return status, err
	}

	return status, nil
}

// @Summary : Insert status
// @Description : insert status to repository
// @Author : rasmadibnu
func (s *StatusService) Insert(req request.Status) (entity.Status, error) {
	status := entity.Status{
		Name:        req.Name,
		Color:       req.Color,
		Description: req.Description,
	}

	newStatus, err := s.repository.Insert(status)

	if err != nil {
		return newStatus, err
	}

	return newStatus, nil
}

// @Summary : Find status
// @Description : Find status by id from repository
// @Author : rasmadibnu
func (s *StatusService) FindById(ID int) (entity.Status, error) {
	status, err := s.repository.FindById(ID)

	if err != nil {
		return status, err
	}

	return status, nil
}

// @Summary : Update status
// @Description : Update status by id from repository
// @Author : rasmadibnu
func (s *StatusService) Update(req request.Status, ID int) (entity.Status, error) {
	status := entity.Status{
		Name:        req.Name,
		Color:       req.Color,
		Description: req.Description,
	}

	updateStatus, err := s.repository.Update(status, ID)

	if err != nil {
		return updateStatus, err
	}

	return updateStatus, nil
}

// @Summary : Delete status
// @Description : Delete status from repository
// @Author : rasmadibnu
func (s *StatusService) Delete(ID int) (bool, error) {
	status, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return status, nil
}
