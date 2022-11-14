package main

import (
	"log"
	"net/http"
	"os"

	"atypicaldev.com/conversation/notes_fe/routes"
	"atypicaldev.com/conversation/notes_fe/shared"

	pb "atypicaldev.com/conversation/api/notes"
	"github.com/gin-gonic/gin"
)

type clientBuilder func(log.Logger) (pb.NotesServiceClient, closer)

type closer interface {
	Close() error
}

type server struct {
	engine *gin.Engine
}

func createServer(builder clientBuilder) *server {
	return &server{engine: initGinEngine(builder)}
}

func initGinEngine(builder clientBuilder) *gin.Engine {
	engine := gin.Default()

	// create client
	engine.Use(func(ctx *gin.Context) {
		l := log.New(os.Stdout,
			"NotesServiceClientBuilder",
			log.Lmsgprefix|log.Llongfile|log.Ldate|log.Ltime,
		)
		client, closer := builder(*l)
		ctx.Set(shared.NotesServiceKey, client)

		ctx.Next()

		err := closer.Close()

		if err != nil {
			log.Printf("Something bad happened closing connection to notes service: %v", err)
		}
	})

	registerNotes(engine)
	return engine
}

func registerNotes(e *gin.Engine) {
	group := e.Group("/notes")

	group.GET("/", routes.GetNotesList)
}

func (s server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.engine.ServeHTTP(rw, r)
}
