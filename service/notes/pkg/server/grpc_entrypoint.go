package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
	"atypicaldev.com/conversation/notes/pkg/data"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func InitServer(dbUrl, dbEnv, grpcServerPort string) {
	notesRepository := data.NewRepository(dbUrl, dbEnv)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(middleWare()...)
	reflection.Register(grpcServer)

	service_pb.RegisterConversationServiceServer(grpcServer, newServer(notesRepository))
	log.Fatalln(grpcServer.Serve(lis))
}

func InitGatewayServer(grpcServerPort, gatewayPort string) {
	conn, err := grpc.Dial(
		fmt.Sprintf("0.0.0.0:%s", grpcServerPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	err = service_pb.RegisterConversationServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", gatewayPort),
		Handler: cors(gatewayLogger(gwmux)),
	}

	log.Printf("Serving Gateway on http://localhost:%s\n", gatewayPort)
	log.Fatalln(gwServer.ListenAndServe())
}
