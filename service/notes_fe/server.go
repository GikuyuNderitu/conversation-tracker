package main

import (
	"net/http"

	"atypicaldev.com/conversation/notes_fe/middleware"
	"atypicaldev.com/conversation/notes_fe/routes"

	"github.com/gin-gonic/gin"
)

type server struct {
	engine *gin.Engine
}

func createServer(builder middleware.ClientBuilder) *server {
	return &server{engine: initGinEngine(builder)}
}

func initGinEngine(builder middleware.ClientBuilder) *gin.Engine {
	engine := gin.Default()

	// Middleware
	engine.Use(middleware.RegisterRouteLogger)
	engine.Use(middleware.RegisterNotesServiceClient(builder))

	// Route groups
	registerNotes(engine)
	registerConvos(engine)

	return engine
}

func registerNotes(e *gin.Engine) {
	group := e.Group("/notes")

	group.GET("/", routes.GetNotesList)
}

func registerConvos(e *gin.Engine) {
	group := e.Group("/convos")

	group.GET("/", routes.GetConversations)
	group.GET("/:id", routes.GetConversationDetail)
	group.POST("/", routes.CreateConversation)
}

func (s server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.engine.ServeHTTP(rw, r)
}
