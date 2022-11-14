package middleware

import (
	"log"

	pb "atypicaldev.com/conversation/api/notes"
	"atypicaldev.com/conversation/notes_fe/shared"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ClientBuilder func(*logrus.Logger) (pb.NotesServiceClient, Closer)

type Closer interface {
	Close() error
}

// Populates context with a the client to call the notes service
func RegisterNotesServiceClient(builder ClientBuilder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l := ctx.Value(shared.LoggerKey).(*logrus.Logger)
		client, closer := builder(l)
		ctx.Set(shared.NotesServiceKey, client)

		ctx.Next()

		err := closer.Close()

		if err != nil {
			log.Printf("Something bad happened closing connection to notes service: %v", err)
		}
	}
}
