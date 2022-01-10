package cpm

type CpmImport struct {
	ProjectId int    `json:"projectId"  form:"project_id" gorm:"comment:项目id"`
	Version   string `json:"version" gorm:"comment:版本号"`
	Name      string `json:"name" gorm:"comment:组件名字"`
	Url       string `json:"url" gorm:"comment:js路径"`
	Md        string `json:"md" gorm:"comment:js路径"`
}
