package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/sentry"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	pb "github.com/timurkash/kratos-layout/api/helloworld/v1"
	"github.com/timurkash/kratos-layout/internal/conf"
	"github.com/timurkash/kratos-layout/internal/service"
)

// NewHTTPServer new HTTP server.
func NewHTTPServer(confServer *conf.Server, greeter *service.GreeterService, logger log.Logger) (*http.Server, error) {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			sentry.Server(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if confServer.Http.Network != "" {
		opts = append(opts, http.Network(confServer.Http.Network))
	}
	if confServer.Http.Addr != "" {
		opts = append(opts, http.Address(confServer.Http.Addr))
	}
	if confServer.Http.Timeout != nil {
		opts = append(opts, http.Timeout(confServer.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	pb.RegisterGreeterHTTPServer(srv, greeter)
	return srv, nil
}
