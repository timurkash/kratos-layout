package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/sentry"
	pb "github.com/timurkash/kratos-layout/api/helloworld/v1"
	"github.com/timurkash/kratos-layout/internal/conf"
	"github.com/timurkash/kratos-layout/internal/service"
	"time"
)

type Jwks struct {
	Url              string
	RefreshInterval  time.Duration
	RefreshRateLimit time.Duration
	RefreshTimeout   time.Duration
}

func authMiddleware(jwks *Jwks) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			return
		}
	}
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(confServer *conf.Server, confJwks *conf.Jwks, greeter *service.GreeterService, logger log.Logger) (*grpc.Server, error) {
	jwks := &Jwks{}
	if confJwks != nil {
		jwks.Url = confJwks.Url
		jwks.RefreshRateLimit = confJwks.RefreshRateLimit.AsDuration()
		jwks.RefreshInterval = confJwks.RefreshInterval.AsDuration()
		jwks.RefreshTimeout = confJwks.RefreshTimeout.AsDuration()
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			sentry.Server(),
			tracing.Server(),
			logging.Server(logger),
			authMiddleware(jwks),
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
	pb.RegisterGreeterServer(srv, greeter)
	return srv, nil
}
