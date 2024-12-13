package repository

import (
	"gorbac/app/entity"
	"gorbac/config"

	"gorm.io/gorm"
)

type ProductsRepository struct {
	config config.Database
}

func NewProductsRepository(db config.Database) ProductsRepository {
	return ProductsRepository{
		config: db,
	}
}

func (r *ProductsRepository) Index() *gorm.DB {
	return r.config.DB.Model(&entity.Products{})
}
