// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mongo

import (
	"context"
	"fmt"
	log "github.com/kingwel-xie/k2/core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"wallet-example/config"
	mongodb "wallet-example/internal/pkg/mongdb"
	"wallet-example/internal/wallet/store"
)

type datastore struct {
	db *mongo.Client

	// can include two database instance if needed
	// docker *grom.DB
	// db *gorm.DB
}

func (ds *datastore) Close() error {
	return ds.db.Disconnect(context.TODO())
}

var (
	mongoFactory store.MongoFactory
	once         sync.Once
)

// GetMongoFactoryOr create mysql factory with the given config.
func GetMongoFactoryOr() (store.MongoFactory, error) {
	if mongoFactory != nil {
		return mongoFactory, nil
	}
	opts := config.Extend.MongoDB
	options := &mongodb.Options{
		Hosts:                   opts.Hosts,
		UserName:                opts.UserName,
		Password:                opts.Password,
		MaxPoolSize:             opts.MaxPoolSize,
		DbName:                  opts.DbName,
		ReplicaSet:              opts.ReplicaSet,
		ReadPreference:          opts.ReadPreference,
		ServerSelectionTimeoutS: opts.ServerSelectionTimeoutS,
	}
	dbIns, err := mongodb.New(options)
	if err != nil {
		log.Fatalf("failed to get mongo store fatory, mongoFactory : %v", err)
		return nil, fmt.Errorf("failed to get mongo store fatory, mongoFactory: %+v, error: %w", mongoFactory, err)
	}
	mongoFactory = &datastore{dbIns}
	log.Info("mongo db factory init success!")
	return mongoFactory, nil
}
