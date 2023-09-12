package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/timurkash/kratos-layout/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the service
	Name string
	// Version is the version of the service
	Version string
	// flagConf is the config flag.
	flagConf string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.StringVar(&flagConf, "conf", "./configs", "config path, eg: -conf config.yaml")
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	theConfig := config.New(
		config.WithSource(
			file.NewSource(flagConf),
			env.NewSource(""),
		),
	)
	if err := theConfig.Load(); err != nil {
		panic(err)
	}
	var bootstrap conf.Bootstrap
	if err := theConfig.Scan(&bootstrap); err != nil {
		panic(err)
	}
	if err := theConfig.Close(); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bootstrap.Server, bootstrap.Business, bootstrap.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
