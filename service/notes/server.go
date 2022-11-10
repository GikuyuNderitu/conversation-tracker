package main

import (
	"context"

	pb "atypicaldev.com/conversation/notes/api"
)

type conversationServer struct {
	pb.UnimplementedNotesServiceServer
}

func (s *conversationServer) GetNotes(ctx context.Context, request *pb.GetNotesRequest) (response *pb.GetNotesResponse, err error) {
	return
}
func (s *conversationServer) GetConversation(ctx context.Context, request *pb.GetConversationRequest) (response *pb.GetConversationResponse, err error) {
	return
}
func (s *conversationServer) ListConversations(ctx context.Context, request *pb.ListConversationsRequest) (response *pb.ListConversationsResponse, err error) {
	return
}

func newConversationServer() *conversationServer {
	return &conversationServer{}
}
