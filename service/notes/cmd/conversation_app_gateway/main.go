package main

import (
	"atypicaldev.com/conversation/notes/pkg/server"
)

// Initialize and start the notes service
func main() {
	serverPort := "8090"
	gatewayPort := "1337"

	server.InitGatewayServer(serverPort, gatewayPort)
}
