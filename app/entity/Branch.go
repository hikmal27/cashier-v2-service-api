package entity

type Branch struct {
	ID          int        `gorm:"primary_key; column:branchID"`
	FullName    string     `gorm:"column:branchFullName"`
	InitName    string     `gorm:"column:branchInitialName"`
	AreaID      int        `gorm:"column:areaID"`
	BranchArea  BranchArea `gorm:"foreignKey:AreaID;references:AreaID"`
	Address     string     `gorm:"size:255; column:branchAddress"`
	Province    string     `gorm:"size:255; column:branchProvince"`
	Regencies   string     `gorm:"size:255; column:branchRegencies"`
	District    string     `gorm:"size:255; column:branchDistrict"`
	SubDistrict string     `gorm:"size:255; column:branchSubDistrict"`
	ZipCode     string     `gorm:"size:255; column:branchZipCode"`
	// Warehouses  []Warehouse `gorm:"many2many:tr_branch_warehouse;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TinyBranch struct {
	ID       int    `gorm:"primary_key; column:branchID"`
	FullName string `gorm:"column:branchFullName"`
	InitName string `gorm:"column:branchInitialName"`
	AreaID   int    `gorm:"column:areaID"`
	// BranchArea  BranchArea  `gorm:"foreignKey:AreaID;references:AreaID"`
	// Address     string      `gorm:"size:255; column:branchAddress"`
	// Province    string      `gorm:"size:255; column:branchProvince"`
	// Regencies   string      `gorm:"size:255; column:branchRegencies"`
	// District    string      `gorm:"size:255; column:branchDistrict"`
	// SubDistrict string      `gorm:"size:255; column:branchSubDistrict"`
	// ZipCode     string      `gorm:"size:255; column:branchZipCode"`
	// Warehouses  []Warehouse `gorm:"many2many:tr_branch_warehouse;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TablerBranch interface {
	TableName() string
}

func (TinyBranch) TableName() string {
	return "ms_branch"
}

func (Branch) TableName() string {
	return "ms_branch"
}
