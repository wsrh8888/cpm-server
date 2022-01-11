package cpm

import "cpm/model/system"

type CpmVersion struct {
	ID          uint   `gorm:"primary_key"` // 主键ID
	ProjectId   int    `json:"projectId" gorm:"comment:'项目ID;version'"`
	Version     string `json:"version" gorm:"comment:'项目ID;version'"`
	PublishId   int    `json:"publisherId"  gorm:"comment:'发布者'"`
	Description string `json:"description"  gorm:"comment:'描述信息'"`
	Keywords    string `json:"keywords" gorm:"comment:'关键字'"`
	Publish     system.SysUser
	ImportList  []CpmImport `json:"ImportList"  gorm:"-"`
}
