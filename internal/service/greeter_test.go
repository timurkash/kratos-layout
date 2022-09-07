package service

import (
	"context"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/timurkash/kratos-layout/internal/biz"
	"github.com/timurkash/kratos-layout/internal/outside/data"
	"github.com/timurkash/kratos-layout/internal/test"
)

func getTestService() (*GreeterService, context.Context, error) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	test.IsUnitTest = true
	dataData, _, err := data.NewData(nil, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(nil, greeterRepo, logger)
	return NewGreeterService(greeterUsecase, logger), context.TODO(), nil
}

func Test_SayHello(t *testing.T) {
	service, ctx, err := getTestService()
	if err != nil {
		t.Fatal(err)
	}
	service = service
	ctx = ctx
	//{
	//	in := &v1.HelloRequest{
	//		Name: test.Test,
	//	}
	//	reply, err := service.SayHello(ctx, in)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if reply.Message != "Hello "+test.Test {
	//		t.Fatal("unexpected message")
	//	}
	//}
	//{
	//	in := &v1.HelloRequest{
	//		Name: test.Error,
	//	}
	//	_, err := service.SayHello(ctx, in)
	//	if err == nil {
	//		t.Fatal("expected error")
	//	}
	//}
}
