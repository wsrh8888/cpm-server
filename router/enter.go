package router

import (
	"cpm/router/cpm"
	system "cpm/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Cpm    cpm.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
