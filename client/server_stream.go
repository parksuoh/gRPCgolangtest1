package main

import (
	"log"
	"io"
	"context"

	pb "github.com/parksuoh/basic-go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("스트리밍 시작")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("이름을 전송할수없음 %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("스트리밍중 에러남 %v", err)
		}
		log.Println(message)
	}
	log.Println("스트리밍 끝남")
}