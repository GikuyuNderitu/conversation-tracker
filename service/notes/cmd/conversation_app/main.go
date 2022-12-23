package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"atypicaldev.com/conversation/notes/pkg/server"
	"github.com/joho/godotenv"
)

const (
	surrealDbUrlKey = "SURREAL_DB_URL"
	surrealDbEnvKey = "SURREAL_DB_ENV"

	serverPortKey    = "CONVO_GRPC_PORT"
	postgresConvoKey = "CONVO_POSTGRES_URL"
)

// Initialize and start the notes service
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
	surrealDbUrl := os.Getenv(surrealDbUrlKey)
	surrealDbEnv := os.Getenv(surrealDbEnvKey)
	serverPort := os.Getenv(serverPortKey)
	postgres := os.Getenv(postgresConvoKey)
	usePostgres := flag.Bool("usepostgres", false, "If true, uses postgres as the data layer")
	flag.Parse()
	fmt.Printf("Value of flag: %v\n", *usePostgres)
	config := server.ServerConfig{
		ServerPort:   serverPort,
		SurrealDBUrl: surrealDbUrl,
		SurrealDBEnv: surrealDbEnv,
		PsqlUrl:      postgres,
		UsePostgres:  usePostgres,
	}

	server.InitServer(config)
}
