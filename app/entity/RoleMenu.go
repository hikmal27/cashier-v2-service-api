package entity

import "time"

type RoleMenu struct {
	ID           int `gorm:"primary_key; not null"`
	RoleID       int `gorm:"not null" json:"role_id"`
	MenuID       int `gorm:"not null" json:"menu_id"`
	CreatedAt    time.Time
	CreateUserID int  `gorm:"not null; column:created_by" json:"created_by"`
	CreatedBy    User `gorm:"foreignKey:CreateUserID"`
}

type TablerRoleMenu interface {
	TableName() string
}

func (RoleMenu) TableName() string {
	return "tr_role_menus"
}
