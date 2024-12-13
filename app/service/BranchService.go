package service

import (
	"gorbac/app/entity"
	"gorbac/app/repository"
)

type BranchService struct {
	repository repository.BranchRepository
}

func NewBranchService(r repository.BranchRepository) BranchService {
	return BranchService{
		repository: r,
	}
}

// @Summary : Find branch
// @Description : Find all branch rom repository
// @Author : rasmadibnu
func (s *BranchService) FindAll(param map[string]interface{}) ([]entity.Branch, error) {
	branch, err := s.repository.FindAll(param)

	if err != nil {
		return branch, err
	}

	return branch, err
}

// @Summary : Find branch
// @Description : Find branch by id from repository
// @Author : rasmadibnu
func (s *BranchService) FindById(ID int) (entity.Branch, error) {
	branch, err := s.repository.FindById(ID)

	if err != nil {
		return branch, err
	}

	return branch, nil
}
