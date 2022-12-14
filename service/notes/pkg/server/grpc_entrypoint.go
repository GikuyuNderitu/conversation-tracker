package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
	"atypicaldev.com/conversation/notes/pkg/data"
	"atypicaldev.com/conversation/notes/pkg/data/postgres"
	"atypicaldev.com/conversation/notes/pkg/data/surreal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type ServerConfig struct {
	SurrealDBUrl, SurrealDBEnv string
	PsqlUrl                    string
	ServerPort                 string
	UsePostgres                *bool
}

func InitServer(config ServerConfig) {
	if config.UsePostgres == nil {
		*config.UsePostgres = false
	}

	notesRepository := getRepository(config)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.ServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(middleWare()...)
	reflection.Register(grpcServer)

	service_pb.RegisterConversationServiceServer(grpcServer, newServer(notesRepository))
	log.Fatalln(grpcServer.Serve(lis))
}

func getRepository(config ServerConfig) data.NotesRepository {
	if *config.UsePostgres {
		fmt.Println("Using postgres as datalayer")
		return postgres.NewPsqlRepository(config.PsqlUrl)
	}

	fmt.Println("Using surrealDB as datalayer")
	return surreal.NewSurrealRepository(config.SurrealDBUrl, config.SurrealDBEnv)
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

func MigratePostgres(config ServerConfig) {
	if config.PsqlUrl == "" {
		log.Fatal("Require the postgres connection url to be supplied")
	}
	postgres.MigratePostgres(config.PsqlUrl)
}
