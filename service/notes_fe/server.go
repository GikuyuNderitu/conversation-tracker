package main

import (
	"log"
	"net/http"
	"os"

	"atypicaldev.com/conversation/notes_fe/routes"
	"atypicaldev.com/conversation/notes_fe/shared"

	pb "atypicaldev.com/conversation/api/notes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	// Middleware
	engine.Use(registerNotesServiceClient(builder))
	engine.Use(registerRouteLogger)

	// Route groups
	registerNotes(engine)
	registerConvos(engine)

	return engine
}

// Populates context with a the client to call the notes service
func registerNotesServiceClient(builder clientBuilder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
	}
}
func registerRouteLogger(ctx *gin.Context) {
	logger := logrus.New()

	ctx.Set(shared.LoggerKey, logger)
	ctx.Next()
}

func registerNotes(e *gin.Engine) {
	group := e.Group("/notes")

	group.GET("/", routes.GetNotesList)
}

func registerConvos(e *gin.Engine) {
	group := e.Group("/convos")

	group.GET("/", routes.GetConversations)
	group.PATCH("/", routes.CreateConversation)
}

func (s server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.engine.ServeHTTP(rw, r)
}
