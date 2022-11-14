package main

import (
	"log"
	"net/http"

	pb "atypicaldev.com/conversation/api/notes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := ":1337"
	http.HandleFunc("/notes", getNotes)
	log.Fatalf("Server unhealthy: %v", http.ListenAndServe(address, createServer(dialService)))
}

func getNotes(rw http.ResponseWriter, r *http.Request) {

}

func dialService(logger log.Logger) (pb.NotesServiceClient, closer) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	address := "localhost:8090"
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		logger.Fatalf("Error dialog NotesService")
	}
	return pb.NewNotesServiceClient(conn), conn
}
