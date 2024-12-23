package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/sentry"
	"github.com/google/wire"
	pb "github.com/timurkash/kratos-layout/gen/go/api/helloworld/v1"
	"github.com/timurkash/kratos-layout/internal/conf"
	"github.com/timurkash/kratos-layout/internal/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(confServer *conf.Server, service *service.GreeterService, logger log.Logger) (*grpc.Server, error) {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			sentry.Server(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	if confServer.Grpc.Network != "" {
		opts = append(opts, grpc.Network(confServer.Grpc.Network))
	}
	if confServer.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(confServer.Grpc.Addr))
	}
	if confServer.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(confServer.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(srv, service)
	return srv, nil
}

// NewHTTPServer new HTTP server.
func NewHTTPServer(confServer *conf.Server, service *service.GreeterService, logger log.Logger) (*http.Server, error) {
	return new(http.Server), nil
}
