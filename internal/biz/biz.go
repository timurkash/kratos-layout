package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/timurkash/kratos-layout/internal/conf"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	CreateGreeter(context.Context, *Greeter) error
	UpdateGreeter(context.Context, *Greeter) error
}

type GreeterUsecase struct {
	confBusiness *conf.Business
	repo         GreeterRepo
	log          *log.Helper
}

func NewGreeterUsecase(confBusiness *conf.Business, repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{
		confBusiness: confBusiness,
		repo:         repo,
		log:          log.NewHelper(logger),
	}
}

func (uc *GreeterUsecase) Create(ctx context.Context, g *Greeter) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GreeterUsecase) Update(ctx context.Context, g *Greeter) error {
	return uc.repo.UpdateGreeter(ctx, g)
}
