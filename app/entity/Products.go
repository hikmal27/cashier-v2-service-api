package entity

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID          int     `gorm:"primary_key; not null"`
	Title       string  `gorm:"size:255; default: NULL"`
	CategoryID  int     `gorm:"default: NULL"`
	Category    string  `gorm:"size:255; default: NULL"`
	Price       float64 `gorm:"type: decimal(10,2); default: NULL"`
	Description string  `gorm:"type: text; default: NULL"`
	Image       string  `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type TablerProducts interface {
	TableName() string
}

func (Products) TableName() string {
	return "ms_products"
}
