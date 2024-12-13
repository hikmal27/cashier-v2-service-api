package request

type Status struct {
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Description *string `json:"description"`
	CreatedBy   string  `json:"created_by"`
	UpdatedBy   string  `json:"updated_by"`
}
