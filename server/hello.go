package server

import (
	"context"
	"fmt"
	pb "grpc-gateway-example/proto/hello"
)

type helloService struct {
	pb.UnimplementedHelloWorldServer
}

func NewHelloService() *helloService {
	return &helloService{}
}

func (h helloService) SayHelloWorld(ctx context.Context, r *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	fmt.Println("成功调用了SayHelloWorld方法了")
	return &pb.HelloWorldResponse{
		Message: "test",
	}, nil
}
