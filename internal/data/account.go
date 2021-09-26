package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lynndotconfig/gkgo2/internal/biz"
)

type accountrRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountrRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountrRepo) CreateAccount(ctx context.Context, a *biz.Account) error {
	return nil
}

func (r *accountrRepo) UpdateAccount(ctx context.Context, a *biz.Account) error {
	return nil
}
