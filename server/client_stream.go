package main

import (
	"log"
	"io"
	pb "github.com/parksuoh/basic-go-grpc/proto"
)

func (s *helloServer) SayhelloClientStreaming(stream pb.GreetService_SayhelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("이름 요청 받음: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}

}
