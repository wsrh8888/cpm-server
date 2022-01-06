package router

import (
	"cpm/router/example"
	system "cpm/router/system"
)

type RouterGroup struct {
	Example example.RouterGroup
	System  system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
