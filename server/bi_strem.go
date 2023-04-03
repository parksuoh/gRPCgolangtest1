package main

import (
	"log"
	"io"
	pb "github.com/parksuoh/basic-go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionsalStreaming(stream pb.GreetService_SayHelloBidirectionsalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("이름받음 %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}