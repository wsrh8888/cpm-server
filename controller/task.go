package controller

import (
	"cpm/initialize"
	"cpm/model"
	"cpm/verification"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveTask(ctx *gin.Context) {
	json := verification.SaveTask{}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 1, "mgs": err.Error()})
		return
	}
	fromt_task := model.Task{
		Version: json.Version,
		Name:    json.Name,
		Main:    json.Main,
		Author:  json.Author,
	}
	if err := initialize.DB.Create(&fromt_task); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "保存成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
	}
}
