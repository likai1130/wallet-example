package blockchain

import (
	"github.com/gin-gonic/gin"
	log "github.com/kingwel-xie/k2/core/logger"
	"strconv"
	"wallet-example/internal/pkg/code"
	"wallet-example/internal/pkg/core"
	"wallet-example/internal/pkg/errors"
)

// Delete 删除一个链
// @Summary 根据id删除区块链
// @Description 根据id删除区块链
// @Tags 区块链管理
// @Param id path int true "主键Id"
// @Success 200 {object} core.Response{} "{"code": 200, message:"Ok"}"
// @Success 400 {object} core.Response
// @Success 404 {object} core.Response
// @Success 500 {object} core.Response
// @Router /v1/chains/{id} [delete]
func (u *BlockchainController) Delete(c *gin.Context) {
	log.Info("delete chain function called.")
	param := c.Param("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, err.Error()), nil)
	}

	if err = u.srv.Blockchains().Delete(c, id); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
