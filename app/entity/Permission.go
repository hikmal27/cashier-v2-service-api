package entity

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          int    `gprm:"primary_key; not null"`
	Name        string `gorm:"size:100; not null"`
	Description string `gorm:"default: NULL"`
	CreatedAt   time.Time
	CreatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt   time.Time
	UpdatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt   gorm.DeletedAt
}

type TablerPermission interface {
	TableName() string
}

func (Permission) TableName() string {
	return "ms_permissions"
}
