package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Account struct {
	Name string
	Email string
}

type AccountRepo interface {
	CreateAccount(context.Context, *Account) error
	UpdateAccount(context.Context, *Account) error
}

type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUsecase) Create(ctx context.Context, g *Account) error {
	return uc.repo.CreateAccount(ctx, g)
}

func (uc *AccountUsecase) Update(ctx context.Context, g *Account) error {
	return uc.repo.UpdateAccount(ctx, g)
}
