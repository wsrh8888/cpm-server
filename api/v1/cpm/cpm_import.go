package cpm

import (
	"cpm/model/cpm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*VersionApi) AddCpmImport(c *gin.Context, list []cpm.CpmImport) {

	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"result": list,
	})
	// var cpmImport []cpm.CpmImport

	// if info, err := cpmService.AddImport(cpmImport); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code":   0,
	// 		"result": info,
	// 	})
	// }
}
