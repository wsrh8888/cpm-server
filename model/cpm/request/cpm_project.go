package request

import (
	"cpm/model/common/request"
	"cpm/model/cpm"
)

type SysDictionaryDetailSearch struct {
	cpm.CpmProject
	request.PageInfo
}
