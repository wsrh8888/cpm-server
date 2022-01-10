package cpm

type CpmProjectType struct {
	TypeId   string `json:"typeId" gorm:"comment:语言ID"`
	TypeName string `json:"typeName" gorm:"comment:语言名称"`
}
