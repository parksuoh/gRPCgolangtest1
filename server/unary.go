package main

import (
	"context"
	pb "github.com/parksuoh/basic-go-grpc/proto"
)

func(s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}





