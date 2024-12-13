package service

import (
	"gorbac/app/repository"

	"gorm.io/gorm"
)

type ProductsService struct {
	repository repository.ProductsRepository
}

func NewProductsService(r repository.ProductsRepository) ProductsService {
	return ProductsService{
		repository: r,
	}
}

func (s *ProductsService) FindAll() *gorm.DB {
	return s.repository.Index()
}
