package request

import (
	"cpm/model/cpm"
)

type CpmVersionAdd struct {
	cpm.CpmVersion
	CpmImportList []cpm.CpmImport `json:"importList" gorm:"-"`
}
