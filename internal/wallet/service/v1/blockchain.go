package v1

import (
	"context"
	log "github.com/kingwel-xie/k2/core/logger"
	"regexp"
	"wallet-example/internal/pkg/code"
	"wallet-example/internal/pkg/errors"
	metav1 "wallet-example/internal/wallet/meta/v1"
	v1 "wallet-example/internal/wallet/models/v1"
	"wallet-example/internal/wallet/service/v1/dto"
	"wallet-example/internal/wallet/store"
)

type BlockchainSrv interface {
	Create(ctx context.Context, chain *dto.BlockchainInsertReq) error
	Update(ctx context.Context, id int64, chain *dto.BlockchainUpdateReq) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*v1.BlockChain, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BlockChainList, error)
}

type blockchainService struct {
	store store.MysqlFactory
}

var _ BlockchainSrv = (*blockchainService)(nil)

func newBlockchains(srv *service) *blockchainService {
	return &blockchainService{store: srv.mysqlStore}
}

// List returns user list in the storage. This function has a good performance.
func (u *blockchainService) List(ctx context.Context, opts metav1.ListOptions) (*v1.BlockChainList, error) {
	blockchains, err := u.store.BlockChains().List(ctx, opts)
	if err != nil {
		log.Errorf("list users from storage failed: %s", err.Error())

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	log.Debugf("get %d chains from backend storage.", len(blockchains.Items))

	return blockchains, nil
}

func (u *blockchainService) Create(ctx context.Context, chain *dto.BlockchainInsertReq) error {
	var blockchain = &v1.BlockChain{}
	chain.Generate(blockchain)
	if err := u.store.BlockChains().Create(ctx, blockchain); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'chain_id'", err.Error()); match {
			return errors.WithCode(code.ErrBlockChainAlreadyExist, err.Error())
		}

		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (u *blockchainService) Delete(ctx context.Context, id int64) error {
	if err := u.store.BlockChains().Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (u *blockchainService) Get(ctx context.Context, id int64) (*v1.BlockChain, error) {
	chain, err := u.store.BlockChains().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return chain, nil
}

func (u *blockchainService) Update(ctx context.Context, id int64, chain *dto.BlockchainUpdateReq) error {
	blockChain, err := u.store.BlockChains().Get(ctx, id)
	if err != nil {
		return err
	}
	chain.Generate(blockChain)

	if err := u.store.BlockChains().Update(ctx, blockChain); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}
