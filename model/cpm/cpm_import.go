package cpm

type CpmImport struct {
	ID        uint   `gorm:"primary_key"` // 主键ID
	ProjectId int    `json:"projectId" gorm:"comment:项目id"`
	Version   string `json:"version" gorm:"comment:版本号"`
	Name      string `json:"name" gorm:"comment:组件名字"`
	List      string `json:"List" gorm:"comment:js路径"`
	Md        string `json:"md" gorm:"comment:js路径"`
}
