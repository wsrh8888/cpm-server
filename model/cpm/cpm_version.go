package cpm

type CpmVersion struct {
	ProjectId   int    `json:"projectId"  form:"project_id" gorm:"comment:项目id"`
	Version     string `json:"version" gorm:"comment:'版本号'"`
	Publisher   string `json:"publisher"  form:"user_id" gorm:"comment:'发布者'"`
	Description string `json:"description"  gorm:"comment:'描述信息'"`
	Keywords    string `json:"keywords" gorm:"comment:'关键字'"`
}
