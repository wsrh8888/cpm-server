package cpm

import "cpm/service"

type ApiGroup struct {
	ProjectApi
}

var (
	cpmService = service.ServiceGroupApp.CpmServiceGroup.CpmService
)