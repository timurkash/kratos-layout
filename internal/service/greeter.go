package service

import (
	"context"

	pb "github.com/go-kratos/kratos-layout/api/helloworld/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

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
