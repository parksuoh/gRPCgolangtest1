package main

import (
	"log"
	"context"
	"time"
	pb "github.com/parksuoh/basic-go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("할수없음 %v", err)
	}

	log.Printf("%s", res.Message)


}