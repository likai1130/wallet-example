package mysql

import (
	"context"
	"gorm.io/gorm"
	"wallet-example/internal/pkg/code"
	"wallet-example/internal/pkg/errors"
	metav1 "wallet-example/internal/wallet/meta/v1"
	v1 "wallet-example/internal/wallet/models/v1"
	"wallet-example/internal/wallet/store"
)

type blockchains struct {
	db *gorm.DB
}

func newBlockchains(ds *datastore) store.BlockChainStore {
	return &blockchains{ds.db}
}

// Create creates a new chain account.
func (u *blockchains) Create(ctx context.Context, chain *v1.BlockChain) error {
	return u.db.Create(&chain).Error
}

// Update updates an chain account information.
func (u *blockchains) Update(ctx context.Context, chain *v1.BlockChain) error {
	return u.db.Save(chain).Error
}

// Delete deletes the chain by the chain identifier.
func (u *blockchains) Delete(ctx context.Context, id int64) error {
	err := u.db.Unscoped().Where("id = ?", id).Delete(&v1.BlockChain{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Get return an chain by the chain identifier.
func (u *blockchains) Get(ctx context.Context, id int64) (*v1.BlockChain, error) {
	chain := &v1.BlockChain{}
	err := u.db.Where("id = ?", id).First(&chain).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrBlockChainNotFound, err.Error())
		}
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return chain, nil
}

// List return all blockchains.
func (u *blockchains) List(ctx context.Context, opts metav1.ListOptions) (*v1.BlockChainList, error) {
	ret := &v1.BlockChainList{}
	d := u.db.Offset(int(opts.GetPageIndex())).
		Limit(int(opts.GetPageSize())).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)
	ret.SetTotalPage(opts.GetPageSize())
	ret.SetCurrentPage(opts.GetPageIndex())
	return ret, d.Error
}
