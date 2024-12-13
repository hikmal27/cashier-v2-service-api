package entity

import (
	"time"
)

type UserRole struct {
	ID        int `gorm:"primary_key; not null" json:"id"`
	UserID    int `gorm:"not null; index" json:"user_id"`
	RoleID    int `gorm:"not null; index" json:"role_id"`
	CreatedAt time.Time
	CreatedBy int `gorm:"not null;" json:"created_by"`
}

type TablerUserRole interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "tr_user_roles"
}
