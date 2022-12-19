package main

import (
	"atypicaldev.com/conversation/notes/pkg/server"
)

// Initialize and start the notes service
func main() {
	dbUrl := "ws://localhost:9021/rpc"
	dbEnv := "test-dev"
	serverPort := "8090"

	server.InitServer(dbUrl, dbEnv, serverPort)
}
