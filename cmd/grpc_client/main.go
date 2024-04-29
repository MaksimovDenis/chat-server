package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	desc "github.com/MaksimovDenis/chat-server/pkg/chatAPI_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

var t = time.Now()

type server struct {
	desc.UnimplementedChatAPIV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	usernamesInfo := req.GetUsernames().Usernames

	newUsername := &desc.ChatInfo{
		Usernames: usernamesInfo,
	}

	return &desc.CreateResponse{
		Id: newUsername.GetId(),
	}, nil

}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	chatID := req.GetId()

	log.Printf("Delete chat with ID: %v", chatID)
	return &empty.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {

	newMessage := &desc.Chat{
		From:      gofakeit.BeerName(),
		Text:      gofakeit.Company(),
		Timestamp: timestamppb.Now(),
	}

	log.Printf("Message %v", newMessage)

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatAPIV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
