// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import "wallet-example/internal/wallet/store"

//go:generate mockgen -self_package=github.com/marmotedu/iam/internal/apiserver/service/v1 -destination mock_service.go -package v1 github.com/marmotedu/iam/internal/apiserver/service/v1 Service,UserSrv,SecretSrv,PolicySrv

// Service defines functions used to return resource interface.
type Service interface {
	Blockchains() BlockchainSrv
}

type service struct {
	mysqlStore store.MysqlFactory
	mongoStore store.MongoFactory
}

// NewService returns Service interface.
func NewService(mysqlStore store.MysqlFactory, mongoStore store.MongoFactory) Service {
	return &service{
		mysqlStore: mysqlStore,
		mongoStore: mongoStore,
	}
}

func (s *service) Blockchains() BlockchainSrv {
	return newBlockchains(s)
}
