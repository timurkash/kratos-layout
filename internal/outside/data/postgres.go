package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/timurkash/kratos-layout/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo is constructor for greeterRepo
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, greeter *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, greeter *biz.Greeter) error {
	return nil
}
