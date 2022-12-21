package main

import (
	"flag"
	"fmt"

	"atypicaldev.com/conversation/notes/pkg/server"
)

// Initialize and start the notes service
func main() {
	dbUrl := "ws://localhost:9021/rpc"
	dbEnv := "test-dev"
	serverPort := "8090"
	postgres := "postgresql://postgres:WDFbDea0MJBgEdqUXkjq@containers-us-west-128.railway.app:5573/railway"
	usePostgres := flag.Bool("usepostgres", false, "If true, uses postgres as the data layer")
	flag.Parse()
	fmt.Printf("Value of flag: %v\n", *usePostgres)
	config := server.ServerConfig{
		ServerPort:   serverPort,
		SurrealDBUrl: dbUrl,
		SurrealDBEnv: dbEnv,
		PsqlUrl:      postgres,
		UsePostgres:  usePostgres,
	}

	server.InitServer(config)
}
