package biz

import (
	"context"
	"github.com/timurkash/kratos-layout/internal/conf"

	"github.com/MicahParks/keyfunc"
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
	jwks *keyfunc.JWKS
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, confJwks *conf.Jwks, logger log.Logger) (*GreeterUsecase, error) {
	logHelper := log.NewHelper(logger)
	jwks, err := keyfunc.Get(confJwks.Url, keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			logHelper.Errorf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   confJwks.RefreshInterval.AsDuration(),
		RefreshRateLimit:  confJwks.RefreshRateLimit.AsDuration(),
		RefreshTimeout:    confJwks.RefreshTimeout.AsDuration(),
		RefreshUnknownKID: true,
	})
	if err != nil {
		return nil, err
	}
	return &GreeterUsecase{
		repo: repo,
		jwks: jwks,
		log:  logHelper,
	}, nil
}

func (uc *GreeterUsecase) Create(ctx context.Context, g *Greeter) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GreeterUsecase) Update(ctx context.Context, g *Greeter) error {
	return uc.repo.UpdateGreeter(ctx, g)
}
