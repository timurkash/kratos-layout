package biz

import (
	"context"

	"github.com/timurkash/kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	CreateGreeter(context.Context, *Greeter) error
	UpdateGreeter(context.Context, *Greeter) error
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(confBusiness *conf.Business, repo GreeterRepo, logger log.Logger) (*GreeterUsecase, error) {
	return &GreeterUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}, nil
}

func (uc *GreeterUsecase) Create(ctx context.Context, g *Greeter) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GreeterUsecase) Update(ctx context.Context, g *Greeter) error {
	return uc.repo.UpdateGreeter(ctx, g)
}
