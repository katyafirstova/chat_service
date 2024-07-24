package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	fmt.Printf("#%v", req)
	return &chat_v1.CreateResponse{Id: 11}, nil
}

func (s *server) Delete(_ context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("#%v", req)
	return nil, nil
}

func (s *server) Send(_ context.Context, req *chat_v1.SendRequest) (*emptypb.Empty, error) {
	fmt.Printf("#%v", req)
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50001")
	if err != nil {
		log.Fatalf("Failed to create listener: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat_v1.RegisterChatV1Server(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}

}
