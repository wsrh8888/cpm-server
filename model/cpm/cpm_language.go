package cpm

type CpmLanguage struct {
	ID   string `gorm:"primary_key"` // 主键ID
	Name string `json:"languageName" gorm:"comment:语言名称"`
}
