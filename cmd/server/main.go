package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
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

	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bootstrap conf.Bootstrap
	if err := c.Scan(&bootstrap); err != nil {
		panic(err)
	}
	if err := c.Close(); err != nil {
		panic(err)
	}

	//if bootstrap.Trace != nil && bootstrap.Trace.Endpoint != "" {
	//	if err := jaeger.SetTracerProvider(bootstrap.Trace.Endpoint, Name); err != nil {
	//		log.Error(err)
	//	}
	//	log.Info("tracer provider", bootstrap.Trace.Endpoint)
	//}

	//if bootstrap.Sentry != nil && bootstrap.Sentry.Dsn != "" {
	//	if err := sentry.SetSentryDns(bootstrap.Sentry.Dsn); err != nil {
	//		log.Error(err)
	//	}
	//	log.Info("sentry", bootstrap.Sentry.Dsn)
	//}

	app, cleanup, err := initApp(bootstrap.Server, bootstrap.Data, bootstrap.Business, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
