package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID           int     `gprm:"primary_key; not null"`
	Name         string  `gorm:"size:100; not null"`
	Description  *string `gorm:"default: NULL"`
	Menus        []Menu  `gorm:"many2many:tr_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permissions  []RolePermission
	CreatedAt    time.Time
	CreateUserID int  `gorm:"not null; column:created_by"`
	CreatedBy    User `gorm:"foreignKey:CreateUserID"`
	UpdatedAt    time.Time
	UpdateUserID int  `gorm:"not null; column:updated_by"`
	UpdatedBy    User `gorm:"foreignKey:UpdateUserID"`
	DeletedAt    gorm.DeletedAt
}

type TinyRole struct {
	ID          int     `gprm:"primary_key; not null"`
	Name        string  `gorm:"size:100; not null"`
	Description *string `gorm:"default: NULL"`
	// Menus       []TinyMenu       `gorm:"many2many:tr_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Permissions []RolePermission `gorm:"foreignKey:ID"`
}

type TablerRole interface {
	TableName() string
}

func (TinyRole) TableName() string {
	return "ms_roles"
}

func (Role) TableName() string {
	return "ms_roles"
}
