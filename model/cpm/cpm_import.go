package cpm

import uuid "github.com/satori/go.uuid"

type CpmImport struct {
	ID        uint      `gorm:"primary_key"` // 主键ID
	ProjectId uuid.UUID `json:"projectId" gorm:"column:project_id;comment:'项目ID;version'" binding:"required"`
	Version   string    `json:"version" gorm:"comment:版本号" binding:"required"`
	Name      string    `json:"name" gorm:"comment:组件名字" binding:"required"`
	List      string    `json:"List" gorm:"comment:js路径" binding:"required"`
	Md        string    `json:"md" gorm:"comment:js路径"`
}
