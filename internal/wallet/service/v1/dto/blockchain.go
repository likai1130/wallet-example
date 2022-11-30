package dto

import v1 "wallet-example/internal/wallet/models/v1"

type BlockchainInsertReq struct {
	ChainId        uint32 `json:"chain_id"`
	ChainName      string `json:"chain_name"`
	InfraHttp      string `json:"infra_http"`
	InfraWebsocket string `json:"infra_websocket"`
	Symbol         string `json:"symbol"`
}

func (bl *BlockchainInsertReq) Generate(chain *v1.BlockChain) {
	chain.ChainId = bl.ChainId
	chain.InfraHttp = bl.InfraHttp
	chain.InfraWebsocket = bl.InfraWebsocket
	chain.ChainName = bl.ChainName
	chain.Symbol = bl.Symbol
}

type BlockchainUpdateReq struct {
	ChainId        uint32 `json:"chain_id"`
	ChainName      string `json:"chain_name"`
	InfraHttp      string `json:"infra_http"`
	InfraWebsocket string `json:"infra_websocket"`
	Symbol         string `json:"symbol"`
}

func (bl *BlockchainUpdateReq) Generate(chain *v1.BlockChain) {
	chain.ChainId = bl.ChainId
	chain.InfraHttp = bl.InfraHttp
	chain.InfraWebsocket = bl.InfraWebsocket
	chain.ChainName = bl.ChainName
	chain.Symbol = bl.Symbol
}
