package sys

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Version 系统版本
// @Summary 系统版本
// @Description 系统版本查询
// @Tags 系统服务
// @Accept  application/json
// @Product application/json
// @Success 200 {object} Response "{"code": 200, "data":{"version": "latest"} }"
// @Router /version [get]
func Version(c *gin.Context) {
	wd, _ := os.Getwd()
	fileName := "VERSION"
	version := "latest"
	fp := filepath.Join(wd, fileName)

	_, err := os.Open(fp)
	if os.IsNotExist(err) {
		create, _ := os.Create(fp)
		_, _ = create.Write([]byte(version))
	}

	file, err := ioutil.ReadFile(fp)
	if err == nil {
		version = string(file)
	}
	c.JSON(http.StatusOK, &Response{
		Code: 0,
		Data: gin.H{"version": version},
	})
}
