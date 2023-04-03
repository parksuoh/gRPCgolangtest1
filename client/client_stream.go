package main

import (
	"log"
	"context"
	"time"
	pb "github.com/parksuoh/basic-go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("클라이언트 스트리밍 시작")
	stream, err := client.SayhelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("이름을 보낼수 없음 %v", err)
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err !=nil {
			log.Fatalf("전송중 에러남 %v", err)
		}
		log.Printf("요청 보냈음 %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("클라이언트 스트리밍 끝")
	if err != nil {
		log.Fatalf("에러남 %v", err)
	}
	log.Printf("%v", res.Messages)
}
