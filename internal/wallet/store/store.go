package store

type MysqlFactory interface {
	BlockChains() BlockChainStore
}

type MongoFactory interface {
}
