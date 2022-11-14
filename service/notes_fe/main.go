package main

import (
	"log"
	"net/http"

	pb "atypicaldev.com/conversation/api/notes"
	"atypicaldev.com/conversation/notes_fe/middleware"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := ":1337"

	startup()

	log.Fatalf("Server unhealthy: %v", http.ListenAndServe(address, createServer(dialService)))
}

func startup() {
	_, conn := dialService(logrus.New())
	err := conn.Close()
	if err != nil {
		panic("Problem while connecting to service")
	}
}

func dialService(logger *logrus.Logger) (pb.NotesServiceClient, middleware.Closer) {
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
