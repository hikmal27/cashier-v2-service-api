package request

type MenuRequest struct {
	Name      string `json:"name" binding:"required"`
	IsHeader  bool   `json:"is_header"`
	Url       string `json:"url"`
	Icon      string `json:"icon"`
	Ord       int    `json:"index"`
	ParentID  int    `json:"parent_id"`
	CreatedBy int    `json:"created_by"`
	UpdatedBy int    `json:"updated_by"`
}
