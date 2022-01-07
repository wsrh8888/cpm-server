package cpm

type CpmImport struct {
	ProjectId int    `json:"projectId"  form:"project_id" gorm:"comment:项目id"`
	Version   string `json:"version" gorm:"comment:版本号"`
	Name      string `json:"name" gorm:"comment:组件名字"`
	JsUrl     string `json:"jsUrl" gorm:"comment:js路径"`
	CssUrl    string `json:"cssUrl" gorm:"comment:css路径"`
	MdContent string `json:"mdContent" gorm:"comment:md内容"`
}
