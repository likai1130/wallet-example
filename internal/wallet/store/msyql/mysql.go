// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"fmt"
	sdk "github.com/kingwel-xie/k2/common"
	log "github.com/kingwel-xie/k2/core/logger"
	"sync"
	"wallet-example/internal/pkg/errors"
	"wallet-example/internal/wallet/store"

	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB

	// can include two database instance if needed
	// docker *grom.DB
	// db *gorm.DB
}

func (ds *datastore) BlockChains() store.BlockChainStore {
	return newBlockchains(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory store.MysqlFactory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr() (store.MysqlFactory, error) {
	if mysqlFactory != nil {
		return mysqlFactory, nil
	}
	db := sdk.Runtime.GetDb() //从运行时拿db实例
	if db == nil {
		log.Error("GetMySQLFactoryOr sdk runtime get db is nil.")
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, "sdk runtime get db is nil")
	}
	mysqlFactory = &datastore{db}
	return mysqlFactory, nil

}
