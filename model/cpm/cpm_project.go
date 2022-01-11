package cpm

import (
	"cpm/global"
	"cpm/model/system"

	uuid "github.com/satori/go.uuid"
)

type CpmProject struct {
	global.CPM_MODEL
	UUID       uuid.UUID `json:"uuid" gorm:"comment:用户UUID"` // 用户UUID
	Name       string    `json:"name" gorm:"comment:'项目名'"`
	AuthorId   int       `json:"authorId" gorm:"comment:'创建者'"`
	Author     system.SysUser
	TypeId     string `json:"typeId" gorm:"comment:'组件库1/页面2,其他9'"`
	Type       CpmType
	LanguageId string `json:"languageId" gorm:"comment:'语言ID"`
	Language   CpmLanguage
}
