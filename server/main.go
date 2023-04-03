package main

import (
	"log"
	"net"
	pb "github.com/parksuoh/basic-go-grpc/proto"
	"google.golang.org/grpc"
)

const(
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("서버실행에 실패함 %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("서버가 시작됨 %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("실행에 실패함: %v", err)
	}
}