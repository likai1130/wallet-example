package config

import "time"

var Extend Config

// Config 扩展配置
//  extend:
//    demo:
//      name: demo-name
// 使用方法： config.ExtConfig......即可！！
type Config struct {
	SMS     SMS
	MongoDB MongoDBOptions
}

type SMS struct {
	SignName     string `json:"sign-name" mapstructure:"sign-name"`
	TemplateCode string `json:"template-code" mapstructure:"template-code"`
	PublicKey    string `json:"public-key" mapstructure:"public-key"`
	PrivateKey   string `json:"private-key" mapstructure:"private-key"`
	ProjectId    string `json:"project-id" mapstructure:"project-id"`
}

type MongoDBOptions struct {
	Hosts                   []string      `json:"hosts" mapstructure:"hosts"`
	UserName                string        `json:"username" mapstructure:"username"`
	Password                string        `json:"password" mapstructure:"password"`
	MaxPoolSize             uint64        `json:"max-pool-size" mapstructure:"max-pool-size"`
	DbName                  string        `json:"dbname" mapstructure:"dbname"`
	ReplicaSet              string        `json:"replica-set" mapstructure:"replica-set"`
	ReadPreference          string        `json:"read-preference" mapstructure:"read-preference"`
	ServerSelectionTimeoutS time.Duration `json:"server-selection-timeout-s" mapstructure:"server-selection-timeout-s"`
}
