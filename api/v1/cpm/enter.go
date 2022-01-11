package cpm

import "cpm/service"

type ApiGroup struct {
	ProjectApi
	VersionApi
}

var (
	cpmService        = service.ServiceGroupApp.CpmServiceGroup.CpmService
	cpmVersionService = service.ServiceGroupApp.CpmServiceGroup.CpmVersionService
)
