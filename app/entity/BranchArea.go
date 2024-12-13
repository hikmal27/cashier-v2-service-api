package entity

type BranchArea struct {
	ID        int      `gorm:"primary_key;"`
	AreaID    int      `gorm:"column:areaID"`
	FullName  string   `gorm:"column:areaFullName"`
	Active    int      `gorm:"column:isActive"`
	CreatedAt string   `gorm:"column:createdAt"`
	CreatedBy string   `gorm:"column:createdBy"`
	UpdatedAt string   `gorm:"column:updatedAt"`
	UpdatedBy string   `gorm:"column:updatedBy"`
	Branch    []Branch `gorm:"foreignKey:AreaID; references:AreaID"`
}

type TablerBranchArea interface {
	TableName() string
}

func (BranchArea) TableName() string {
	return "ms_area"
}
