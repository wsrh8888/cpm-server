package cpm

import "cpm/model/system"

type CpmVersion struct {
	VersionId   string `json:"versionId" gorm:"primary_key;comment:'项目ID;version'"`
	PublishId   int    `json:"publisherId"  gorm:"comment:'发布者'"`
	Description string `json:"description"  gorm:"comment:'描述信息'"`
	Keywords    string `json:"keywords" gorm:"comment:'关键字'"`
	Publish     system.SysUser
}
