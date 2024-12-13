package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int    `gprm:"primary_key; not null"`
	Username     string `gorm:"size:100; not null"`
	Name         string `gorm:"size:100; not null"`
	Email        string `gorm:"size:100;"`
	Password     string `gorm:"size:100; not null" json:"-"`
	Roles        []Role `gorm:"many2many:tr_user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt    time.Time
	CreateUserID int   `gorm:"not null; column:created_by"`
	CreatedBy    *User `gorm:"foreignKey:CreateUserID"`
	UpdatedAt    time.Time
	UpdateUserID int   `gorm:"not null; column:updated_by"`
	UpdatedBy    *User `gorm:"foreignKey:UpdateUserID"`
	DeletedAt    gorm.DeletedAt
}

type TablerUser interface {
	TableName() string
}

func (User) TableName() string {
	return "ms_users"
}
