package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addres = "localhost:50051"
	userID = 12
)

func main() {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}
	defer conn.Close()

}
