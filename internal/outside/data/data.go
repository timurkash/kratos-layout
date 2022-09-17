package data

import (
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init postgres driver
	_ "github.com/lib/pq"

	"github.com/timurkash/kratos-layout/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	//relational *ent.Client
	log *log.Helper
}

// NewData .
func NewData(confData *conf.Data, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(logger)
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
	driver = driver // TODO
	//if err := cent.SetSchema(driver, confData.Relational.Schema); err != nil {
	//	return nil, nil, err
	//}
	//if confData.Relational.Schema != "" && confData.Relational.Schema != "public" {
	//	if _, err := driver.DB().Exec(
	//		fmt.Sprintf(
	//			"create schema if not exists %s;set search_path to %s;",
	//			confData.Relational.Schema,
	//			confData.Relational.Schema,
	//		),
	//	); err != nil {
	//		return nil, nil, err
	//	}
	//}

	//client := ent.NewClient(ent.Driver(cent.DebugWithContext(driver, logHelper)))

	//sqlDriver := dialect.DebugWithContext(driver, func(ctx context.Context, i ...interface{}) {
	//	logHelper.WithContext(ctx).Info(i...)
	//	_, span := otel.Tracer("entgo.io").Start(ctx,
	//		"Query",
	//		trace.WithAttributes(
	//			attribute.String("sql", fmt.Sprint(i...)),
	//		),
	//		trace.WithSpanKind(trace.SpanKindServer),
	//	)
	//	span.End()
	//})
	//client := ent.NewClient(ent.Driver(sqlDriver))
	//if err := client.Schema.Create(
	//	context.Background(),
	//	migrate.WithDropIndex(true),
	//	migrate.WithDropColumn(true),
	//); err != nil {
	//	logHelper.Errorf("failed creating schema resources: %v", err)
	//	return nil, nil, err
	//}
	data := &Data{
		//relational: client,
		log: logHelper,
	}
	cleanup := func() {
		logHelper.Info("closing the data resources")
		//if err := data.relational.Close(); err != nil {
		//	logHelper.Error(err)
		//}
	}
	return data, cleanup, nil
}
