package main

import (
	"context"
	"log"

	pb "atypicaldev.com/conversation/api/notes"
	"atypicaldev.com/conversation/notes/data"
)

type conversationServer struct {
	pb.UnimplementedNotesServiceServer

	repository data.NotesRepository
}

func (s *conversationServer) GetNote(ctx context.Context, request *pb.GetNoteRequest) (response *pb.GetNoteResponse, err error) {
	response = &pb.GetNoteResponse{Note: s.repository.GetNote(request.NoteId)}
	return
}

func (s *conversationServer) GetNotes(ctx context.Context, request *pb.GetNotesRequest) (response *pb.GetNotesResponse, err error) {
	notes := s.repository.GetNotes()

	log.Printf("Notes from repository: %v", notes)

	response = &pb.GetNotesResponse{
		Notes: notes,
	}
	return
}

func (s *conversationServer) GetConversation(ctx context.Context, request *pb.GetConversationRequest) (response *pb.GetConversationResponse, err error) {
	return
}
func (s *conversationServer) ListConversations(
	ctx context.Context,
	request *pb.ListConversationsRequest,
) (response *pb.ListConversationsResponse, err error) {
	response = &pb.ListConversationsResponse{Conversations: s.repository.ListConversations()}
	return
}

func (s *conversationServer) CreateNote(ctx context.Context, request *pb.CreateNoteRequest) (response *pb.CreateNoteResponse, err error) {
	note := s.repository.CreateNote(request)
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
