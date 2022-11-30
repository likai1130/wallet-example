package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// HealthCheck 健康检查
// @Summary 健康检查
// @Description 系统健康检查
// @Tags 系统服务
// @Accept  application/json
// @Product application/json
// @Success 200 {object} Response "{"code": 200, "data":{"status": "ok"} }"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: 0,
		Data: gin.H{"status": "ok"},
	})
}
