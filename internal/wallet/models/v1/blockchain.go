package v1

import (
	v1 "wallet-example/internal/wallet/meta/v1"
)

type BlockChain struct {
	Id             int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	ChainId        uint32 `json:"chain_id" gorm:"unique;column:chain_id;type:int(32);not null"`
	ChainName      string `json:"chain_name" gorm:"type:varchar(20)"`
	InfraHttp      string `json:"infra_http" gorm:"type:varchar(200)"`
	InfraWebsocket string `json:"infra_websocket" gorm:"type:varchar(200)"`
	Symbol         string `json:"symbol" gorm:"type:varchar(20)"`
	v1.ModelTime
}

type BlockChainList struct {
	v1.ListMeta
	Items []*BlockChain `json:"items"`
}

func (BlockChain) TableName() string {
	return "blockchain"
}
