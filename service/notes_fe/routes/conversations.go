package routes

import (
	"context"

	pb "atypicaldev.com/conversation/api/notes"
	"atypicaldev.com/conversation/notes_fe/shared"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetConversations(ctx *gin.Context) {
	l := ctx.Value(shared.LoggerKey).(logrus.Logger)
	client := ctx.Value(shared.NotesServiceKey).(pb.NotesServiceClient)

	res, err := client.ListConversations(context.Background(), &pb.ListConversationsRequest{})
	if err != nil {
		l.Warningf("Error listing conversations: \n%v", err)
		ctx.Error(err)
		return
	}

	json, err := shared.ToJson(res)
	if err != nil {
		l.Warningf("Error parsing json response:\n%v", err)
	}

	ctx.JSON(200, json)
}

func CreateConversation(ctx *gin.Context) {
	l := ctx.Value(shared.LoggerKey).(logrus.Logger)
	client := ctx.Value(shared.NotesServiceKey).(pb.NotesServiceClient)

	req := &pb.CreateConversationRequest{}
	ctx.ShouldBindJSON(req)
	res, err := client.CreateConversation(context.Background(), req)

	if err != nil {
		l.Warningf("Error creating conversation, \n%v", err)
		return
	}

	json, err := shared.ToJson(res)
	if err != nil {
		l.Warningf("Error parsing json response:\n%v", err)
	}

	ctx.JSON(201, json)
}
