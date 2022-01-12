package cpm

import (
	"cpm/model/common/response"
	"cpm/model/cpm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImportApi struct{}

func (*ImportApi) GetCpmImport(c *gin.Context) {
	var cpmImport cpm.CpmImport
	_ = c.ShouldBindJSON(&cpmImport)
	if info, err := cpmImportService.GetImport(cpmImport); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"result": info,
		})
	}
}

func (*ImportApi) GetCpmImportSearchList(c *gin.Context) {
	var cpmImport []cpm.CpmImport
	_ = c.ShouldBindJSON(&cpmImport)
	if info, err := cpmImportService.GetCpmImportSearchList(cpmImport); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"result": info,
		})
	}
}
