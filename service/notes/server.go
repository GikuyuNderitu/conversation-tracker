package main

import (
	"context"

	pb "atypicaldev.com/conversation/api/notes"
	"atypicaldev.com/conversation/notes/data"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
)

type conversationServer struct {
	pb.UnimplementedNotesServiceServer

	repository data.NotesRepository
}

func (s *conversationServer) GetNote(ctx context.Context, request *pb.GetNoteRequest) (response *pb.GetNoteResponse, err error) {
	note, err := s.repository.GetNote(request.NoteId)
	response = &pb.GetNoteResponse{Note: note}
	return
}

func (s *conversationServer) GetNotes(ctx context.Context, request *pb.GetNotesRequest) (response *pb.GetNotesResponse, err error) {
	notes, err := s.repository.GetNotes()
	l := ctxlogrus.Extract(ctx).Logger

	l.Printf("Notes from repository: %v", notes)

	response = &pb.GetNotesResponse{
		Notes: notes,
	}
	return
}

func (s *conversationServer) GetConversation(ctx context.Context, request *pb.GetConversationRequest) (response *pb.GetConversationResponse, err error) {
	convo, err := s.repository.GetConversation(request.GetConversationId())
	response = &pb.GetConversationResponse{Conversation: convo}
	return
}
func (s *conversationServer) ListConversations(
	ctx context.Context,
	request *pb.ListConversationsRequest,
) (response *pb.ListConversationsResponse, err error) {
	l := ctxlogrus.Extract(ctx).Logger
	convos, err := s.repository.ListConversations()
	l.Infof("Repo Conversations: %v", convos)
	response = &pb.ListConversationsResponse{Conversations: convos}
	return
}

func (s *conversationServer) CreateNote(ctx context.Context, request *pb.CreateNoteRequest) (response *pb.CreateNoteResponse, err error) {
	note, err := s.repository.CreateNote(request)
	response = &pb.CreateNoteResponse{Note: note}
	return
}

func (s *conversationServer) CreateConversation(
	ctx context.Context,
	request *pb.CreateConversationRequest,
) (response *pb.CreateConversationResponse, err error) {
	convo, err := s.repository.CreateConversation(request)
	response = &pb.CreateConversationResponse{Conversation: convo}
	return
}

func newConversationServer(repository data.NotesRepository) *conversationServer {
	return &conversationServer{
		repository: repository,
	}
}
