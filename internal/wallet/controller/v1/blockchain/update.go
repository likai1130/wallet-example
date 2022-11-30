package blockchain

import (
	"github.com/gin-gonic/gin"
	log "github.com/kingwel-xie/k2/core/logger"
	"strconv"
	"wallet-example/internal/pkg/code"
	"wallet-example/internal/pkg/core"
	"wallet-example/internal/pkg/errors"
	"wallet-example/internal/wallet/service/v1/dto"
)

// Update 更新区块链信息
// @Summary 更新区块链信息
// @Description 更新区块链信息
// @Tags 区块链管理
// @Param id path int true "主键id"
// @Param data body dto.BlockchainUpdateReq true "链数据"
// @Success 200 {object} core.Response "{"code": 200, "message":"Ok"}"
// @Success 400 {object} core.Response
// @Success 404 {object} core.Response
// @Success 500 {object} core.Response
// @Router /v1/chains/{id} [put]
func (u *BlockchainController) Update(c *gin.Context) {
	log.Info("update user function called.")

	var r dto.BlockchainUpdateReq

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, err.Error()), nil)
	}

	// Save changed fields.
	if err = u.srv.Blockchains().Update(c, id, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
