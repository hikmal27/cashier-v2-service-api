package entity

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID           int    `gorm:"primary_key; not null"`
	Name         string `gorm:"size:100; not null"`
	Url          string `gorm:"size:100; not null"`
	Icon         string `gorm:"default: NULL"`
	Ord          int    `gorm:"default: NULL"`
	IsHeader     bool   `gorm:"default: 0"`
	ParentID     int    `gorm:"default: NULL"`
	Parent       *Menu  `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Childs       []Menu `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"children"`
	Roles        []Role `gorm:"many2many:tr_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt    time.Time
	CreateUserID int  `gorm:"not null; column:created_by"`
	CreatedBy    User `gorm:"foreignKey:CreateUserID"`
	UpdatedAt    time.Time
	UpdateUserID int  `gorm:"not null; column:updated_by"`
	UpdatedBy    User `gorm:"foreignKey:UpdateUserID"`
	DeletedAt    gorm.DeletedAt
}

type TinyMenu struct {
	ID   int    `gorm:"primary_key; not null"`
	Name string `gorm:"size:100; not null"`
	// Url      string     `gorm:"size:100; not null"`
	// Icon     string     `gorm:"default: NULL"`
	// Ord      int        `gorm:"default: NULL"`
	// IsHeader bool       `gorm:"default: 0"`
	ParentID int        `gorm:"default: NULL"`
	Parent   *Menu      `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Childs   []TinyMenu `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"children"`
	Roles    []TinyRole `gorm:"many2many:tr_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TablerMenu interface {
	TableName() string
}

func (TinyMenu) TableName() string {
	return "ms_menus"
}

func (Menu) TableName() string {
	return "ms_menus"
}
