package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
)

type AreaService struct {
	repository repository.AreaRepository
}

func NewAreaService(r repository.AreaRepository) AreaService {
	return AreaService{
		repository: r,
	}
}

// @Summary : Find area
// @Description : Find all area rom repository
// @Author : rasmadibnu
func (s *AreaService) FindAll() ([]entity.BranchArea, error) {
	areas, err := s.repository.FindAll()

	if err != nil {
		return areas, err
	}

	return areas, err
}

// @Summary : Find area
// @Description : Find area by id from repository
// @Author : rasmadibnu
func (s *AreaService) FindById(ID int) (entity.BranchArea, error) {
	area, err := s.repository.FindById(ID)

	if err != nil {
		return area, err
	}

	return area, nil
}
