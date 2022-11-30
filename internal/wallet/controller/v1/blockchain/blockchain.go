package blockchain

import (
	srvv1 "wallet-example/internal/wallet/service/v1"
	"wallet-example/internal/wallet/store"
)

// BlockchainController create a user handler used to handle request for user resource.
type BlockchainController struct {
	srv srvv1.Service
}

// NewBlockchainController creates a user handler.
func NewBlockchainController(mysqlFactory store.MysqlFactory, mongoFactory store.MongoFactory) *BlockchainController {
	return &BlockchainController{
		srv: srvv1.NewService(mysqlFactory, mongoFactory),
	}
}
