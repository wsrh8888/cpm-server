package cpm

type CpmType struct {
	ID   string `gorm:"primary_key"` // 主键ID
	Name string `json:"Name" gorm:"comment:语言名称"`
}
