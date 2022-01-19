package request

import (
	"cpm/model/common/request"
	"cpm/model/cpm"
)

type SysDictionaryDetailSearch struct {
	cpm.CpmProject
	request.PageInfo
}

type CpmProjectAllInfo struct {
	Name    string `json:"name" gorm:"comment:'项目名'" binding:"required"`
	Version string `json:"version" gorm:"comment:'项目ID;version'"`
}
