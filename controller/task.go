package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SavePackage(ctx *gin.Context) {
	println("dddd")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 505,
		"msg":  "邮箱格式不正确",
	})
}
