package blockchain

import (
	"github.com/gin-gonic/gin"
	log "github.com/kingwel-xie/k2/core/logger"
	"wallet-example/internal/pkg/code"
	"wallet-example/internal/pkg/core"
	"wallet-example/internal/pkg/errors"
	"wallet-example/internal/wallet/service/v1/dto"
)

// Create 创建一个链
// @Summary 创建一个链
// @Description 创建一个链
// @Tags 区块链管理
// @Param data body dto.BlockchainInsertReq true "链数据"
// @Success 200 {object} core.Response{} "{"code": 200, "message":"Ok"}"
// @Success 400 {object} core.Response
// @Success 500 {object} core.Response
// @Router /v1/chains [post]
func (u *BlockchainController) Create(c *gin.Context) {
	log.Info("chain create function called.")
	var r dto.BlockchainInsertReq

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	if err := u.srv.Blockchains().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}
