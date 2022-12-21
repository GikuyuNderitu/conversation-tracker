package main

import (
	"atypicaldev.com/conversation/notes/pkg/server"
)

// Initialize and start the notes service
func main() {
	dbUrl := "ws://localhost:9021/rpc"
	dbEnv := "test-dev"
	serverPort := "8090"
	postgres := "postgresql://postgres:1NSIRE2tfZFfHzsqecZN@containers-us-west-81.railway.app:6874/railway"

	config := server.ServerConfig{
		ServerPort:   serverPort,
		SurrealDBUrl: dbUrl,
		SurrealDBEnv: dbEnv,
		PsqlUrl:      postgres,
	}

	server.InitServer(config)
}
