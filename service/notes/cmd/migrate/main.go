package main

import "atypicaldev.com/conversation/notes/pkg/server"

func main() {
	server.MigratePostgres(server.ServerConfig{
		PsqlUrl: "postgresql://postgres:WDFbDea0MJBgEdqUXkjq@containers-us-west-128.railway.app:5573/railway",
	})
}
