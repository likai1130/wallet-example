package mongodb

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Options struct {
	Hosts                   []string      `json:"hosts"`
	UserName                string        `json:"user-name"`
	Password                string        `json:"password"`
	MaxPoolSize             uint64        `json:"max-pool-size"`
	DbName                  string        `json:"dbname"`
	ReplicaSet              string        `json:"replica-set"`
	ReadPreference          string        `json:"read-preference"`
	ServerSelectionTimeoutS time.Duration `json:"server-selection-timeout-s"`
}

func New(opts *Options) (*mongo.Client, error) {
	ctx := context.Background()
	clientOptions := options.Client()

	clientOptions.SetHosts(opts.Hosts)
	clientOptions.SetMaxPoolSize(opts.MaxPoolSize)                        //最大连接数量
	clientOptions.SetServerSelectionTimeout(opts.ServerSelectionTimeoutS) //连接超时
	clientOptions.SetAppName(opts.DbName)                                 //数据库名
	if opts.UserName != "" && opts.Password != "" {
		clientOptions.SetAuth(options.Credential{Username: opts.UserName, Password: opts.Password})
	}

	clientOptions = clientOptions.SetReplicaSet(opts.ReplicaSet)

	if opts.ReadPreference != "" {
		model, err := readpref.ModeFromString(opts.ReadPreference)
		if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("mongo db set [%s] is error.", opts.ReadPreference))
		}
		pref, err := readpref.New(model)
		if err != nil {
			return nil, errors.WithMessage(err, "mongo db set readPreference is error.")
		}
		clientOptions = clientOptions.SetReadPreference(pref)
	}

	mongoInstance, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.WithMessage(err, "mongo db connect fail.")
	}
	err = mongoInstance.Ping(ctx, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "mongo db ping fail.")
	}
	return mongoInstance, nil
}
