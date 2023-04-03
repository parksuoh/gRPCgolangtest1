package main

import (
	"log"
	"context"
	"time"
	"io"
	pb "github.com/parksuoh/basic-go-grpc/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("양방향 통신 시작")
	stream, err := client.SayHelloBidirectionsalStreaming(context.Background())
	if err != nil {
		log.Fatalf("이름 보낼수 없음 %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil {
				log.Fatalf("스트리밍중 에러남 %v", err)
			}

			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err !=nil {
			log.Fatalf("보내는동안 에러남 %v", err )
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("양방향 통신끝")
}