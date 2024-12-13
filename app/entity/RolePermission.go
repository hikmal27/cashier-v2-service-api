package entity

import (
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	ID           int  `gorm:"primary_key; not null"`
	RoleID       int  `gorm:"not null; index" json:"role_id"`
	MenuID       int  `gorm:"not null; index" json:"menu_id"`
	Create       bool `gorm:"not null" json:"create"`
	Read         bool `gorm:"not null" json:"read"`
	Update       bool `gorm:"not null" json:"update"`
	Delete       bool `gorm:"not null" json:"delete"`
	CreatedAt    time.Time
	CreateUserID int  `gorm:"not null; column:created_by" json:"created_by"`
	CreatedBy    User `gorm:"foreignKey:CreateUserID"`
	UpdatedAt    time.Time
	UpdateUserID int  `gorm:"not null; column:updated_by" json:"updated_by"`
	UpdatedBy    User `gorm:"foreignKey:UpdateUserID"`
	DeletedAt    gorm.DeletedAt
}

type TablerRolePermission interface {
	TableName() string
}

func (RolePermission) TableName() string {
	return "tr_role_permissions"
}
