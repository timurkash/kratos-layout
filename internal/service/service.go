package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	pb "github.com/timurkash/kratos-layout/gen/go/api/helloworld/v1"
	"github.com/timurkash/kratos-layout/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

// GreeterService is a greeter service.
type GreeterService struct {
	pb.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, pb.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
