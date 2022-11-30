package main

import (
	"github.com/kingwel-xie/k2/cmd"
	"wallet-example/cmd/api"
)

//go:generate swag init -o ./api/swagger/docs

// @title wallet API
// @description 重构钱包
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Init("wallet", api.StartCmd)
	cmd.Execute()
}
