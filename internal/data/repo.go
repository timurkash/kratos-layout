package data

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	// init postgres driver
	_ "github.com/lib/pq"

	"github.com/timurkash/kratos-layout/internal/biz"
	"github.com/timurkash/kratos-layout/internal/conf"
)

type greeterRepo struct {
	data *Data
	//relational *ent.Client
}

// NewGreeterRepo is constructor for greeterRepo
func NewGreeterRepo(confData *conf.Data, data *Data) (biz.GreeterRepo, func(), error) {
	if confData.Relational.Dialect != dialect.Postgres {
		return nil, nil, errors.New("not postgres is not supported")
	}
	driver, err := sql.Open(
		dialect.Postgres,
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			confData.Relational.Host,
			confData.Relational.Port,
			confData.Relational.User,
			confData.Relational.Password,
			confData.Relational.Dbname,
			confData.Relational.SslMode,
		),
	)
	if err != nil {
		return nil, nil, err
	}
	_ = driver // TODO
	return &greeterRepo{
			data: data,
		}, func() {
			data.log.Info("closing the data resources")
			//if err := data.relational.Close(); err != nil {
			//	data.log.Error(err)
			//}
		}, nil

}

func (r *greeterRepo) CreateGreeter(ctx context.Context, greeter *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, greeter *biz.Greeter) error {
	return nil
}
