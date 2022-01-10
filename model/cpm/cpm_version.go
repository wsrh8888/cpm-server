package cpm

type CpmVersion struct {
	ProjectId   int    `json:"projectId" gorm:"column:project_id comment:项目id"`
	Version     string `json:"version" gorm:"comment:'版本号'"`
	Publisher   string `json:"publisher"  gorm:"column:user_id comment:'发布者'"`
	Description string `json:"description"  gorm:"comment:'描述信息'"`
	Keywords    string `json:"keywords" gorm:"comment:'关键字'"`
	Main        string `json:"main" gorm:" default:main;comment:'关键字'"`
}
