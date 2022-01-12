package cpm

import "cpm/service"

type ApiGroup struct {
	ProjectApi
	VersionApi
	ImportApi
}

var (
	cpmService        = service.ServiceGroupApp.CpmServiceGroup.CpmService
	cpmVersionService = service.ServiceGroupApp.CpmServiceGroup.CpmVersionService
	cpmImportService  = service.ServiceGroupApp.CpmServiceGroup.CpmImportService
)
