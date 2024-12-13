package entity

type GroupBranch struct {
	ID       int `gorm:"primary_key; not null" json:"id"`
	GroupID  int `gorm:"not null; index" json:"group_id"`
	BranchID int `gorm:"not null; index" json:"branch_id"`
}

type TablerGroupBranch interface {
	TableName() string
}

func (GroupBranch) TableName() string {
	return "tr_group_branch"
}
