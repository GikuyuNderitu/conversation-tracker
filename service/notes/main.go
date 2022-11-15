package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "atypicaldev.com/conversation/api/notes"
	"atypicaldev.com/conversation/notes/data"
)

// Initialize and start the notes service
func main() {
	dbUrl := "ws://localhost:9021/rpc"
	dbEnv := "test-dev"
	notesRepository := data.NewRepository(dbUrl, dbEnv)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8090))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(middleWare()...)
	reflection.Register(grpcServer)

	pb.RegisterNotesServiceServer(grpcServer, newConversationServer(notesRepository))
	grpcServer.Serve(lis)
}
