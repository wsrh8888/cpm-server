package cpm

import (
	"cpm/model/system"
	"time"

	uuid "github.com/satori/go.uuid"
)

type CpmVersion struct {
	ID          uint      `gorm:"primary_key"` // 主键ID
	CreatedAt   time.Time // 创建时间
	ProjectId   uuid.UUID `json:"project_id" gorm:"comment:'项目ID;version'" binding:"required"`
	Version     string    `json:"version" gorm:"comment:'项目ID;version'" binding:"required"`
	PublishId   int       `json:"publisherId"  gorm:"comment:'发布者'" binding:"required"`
	Description string    `json:"description"  gorm:"comment:'描述信息'"`
	Keywords    string    `json:"keywords" gorm:"comment:'关键字'"`
	Url         string    `json:"url" gorm:"comment:'按需引入的域名'"`
	Publish     system.SysUser
	ImportList  []CpmImport `json:"ImportList"  gorm:"-"`
}
