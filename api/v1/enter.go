package v1

import (
	"cpm/api/v1/cpm"
	"cpm/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	CpmApiGroup    cpm.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
