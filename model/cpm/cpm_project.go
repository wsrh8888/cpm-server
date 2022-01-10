package cpm

import (
	"cpm/global"
	"cpm/model/system"

	uuid "github.com/satori/go.uuid"
)

type CpmProject struct {
	global.CPM_MODEL
	UUID     uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"` // 用户UUID
	Name     string         `json:"name" gorm:"comment:'项目名'"`
	Author   int            `json:"author" gorm:"column:user_id;comment:'创建者'"`
	Type     int            `json:"type" gorm:"comment:'组件库1/页面2,其他9'"`
	Language int            `json:"language" gorm:""`
	User     system.SysUser `json:"user"`
}
