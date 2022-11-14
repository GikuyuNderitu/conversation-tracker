package main

import "net/http"

func createServer() *server {
	return &server{}
}

type server struct{}

func (s server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {}
