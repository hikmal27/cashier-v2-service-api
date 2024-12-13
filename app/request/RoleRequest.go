package request

import "gorbac/app/entity"

type RoleRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	CreatedBy   int     `json:"created_by"`
	UpdatedBy   int     `json:"updated_by"`
}

type RoleMenu struct {
	RoleID    int `json:"role_id"`
	MenuID    int `json:"menu_id"`
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
}

type RoleHasMenu struct {
	Menus []entity.RoleMenu `json:"menus"`
}

type RoleHasPermission struct {
	Permission []entity.RolePermission `json:"permissions"`
}
