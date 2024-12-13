package entity

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	ID           int     `gorm:"primary_key; not null"`
	Name         string  `gorm:"size:255; not null"`
	Color        string  `gorm:"size:255; not null"`
	Description  *string `gorm:"size:255; default: null"`
	CreatedAt    time.Time
	CreateUserID int  `gorm:"not null; column:created_by"`
	CreatedBy    User `gorm:"foreignKey:CreateUserID"`
	UpdatedAt    time.Time
	UpdateUserID int  `gorm:"not null; column:updated_by"`
	UpdatedBy    User `gorm:"foreignKey:UpdateUserID"`
	DeletedAt    gorm.DeletedAt
}

type TinyStatus struct {
	ID    int    `gorm:"primary_key; not null"`
	Name  string `gorm:"size:255; not null"`
	Color string `gorm:"size:255; not null"`
}

type TablerStatus interface {
	TableName() string
}

func (TinyStatus) TableName() string {
	return "ms_status"
}

func (Status) TableName() string {
	return "ms_status"
}
