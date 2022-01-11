package cpm

type CpmType struct {
	TypeId   uint   `json:"typeId" gorm:"comment:语言ID;primary_key"`
	TypeName string `json:"typeName" gorm:"comment:语言名称"`
}
