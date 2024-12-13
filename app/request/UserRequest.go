package request

import "gorbac/app/entity"

type RegisterRequest struct {
	// GroupID   int    `json:"group_id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedBy int    `json:"created_by" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	// GroupID   int    `json:"group_id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password"`
	UpdatedBy int    `json:"updated_by"`
}

type UserRole struct {
	UserID    int    `json:"user_id"`
	RoleID    int    `json:"role_id"`
	CreatedBy string `json:"created_by"`
	// UpdatedBy string `json:"updated_by"`
}

type UserHasRole struct {
	Roles []entity.UserRole `json:"roles"`
}

type RefreshToken struct {
	UserID int `json:"user_id" binding:"required"`
	Expire int `json:"expire" binding:"required"`
}
