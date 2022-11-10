package main

import (
	"google.golang.org/grpc"

	pb "atypicaldev.com/conversation/notes/api"
)

// Initialize and start the notes service
func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterNotesServiceServer(grpcServer, newConversationServer())
}
