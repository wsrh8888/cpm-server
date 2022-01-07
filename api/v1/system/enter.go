package system

import "cpm/service"

type ApiGroup struct {
	BaseApi
	DBApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
