package service

import (
	"cpm/service/cpm"
	"cpm/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	CpmServiceGroup    cpm.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
