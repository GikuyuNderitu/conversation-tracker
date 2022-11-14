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
	log.Fatalf("Server unhealthy: %v", http.ListenAndServe(address, nil))
}

func getNotes(rw http.ResponseWriter, r *http.Request) {

}

type server struct {
	client pb.NotesServiceClient
}

func dialService() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	address := "localhost:8090"
	grpc.Dial(address, opts...)
}
