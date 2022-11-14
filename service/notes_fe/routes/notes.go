package routes

import (
	"context"
	"log"

	pb "atypicaldev.com/conversation/api/notes"

	"atypicaldev.com/conversation/notes_fe/shared"
	"github.com/gin-gonic/gin"
)

func GetNotesList(ctx *gin.Context) {
	client := ctx.Value(shared.NotesServiceKey).(pb.NotesServiceClient)

	response, err := client.GetNotes(context.Background(), &pb.GetNotesRequest{})

	if err != nil {
		log.Println("GRPC request failed")
		print(err)
		ctx.Error(err)
	}

	bytes, err := shared.ToJson(response)
	if err != nil {
		ctx.Error(err)
	}
	ctx.JSON(200, string(bytes))
}
