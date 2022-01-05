package core

import (
	"cpm/global"
	"cpm/initialize"
	"fmt"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunWindowsServer() {
	// 初始化全部的路由
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CPM_CONFIG.System.Addr)
	s := initServer(address, Router)
	s.ListenAndServe().Error()
}
