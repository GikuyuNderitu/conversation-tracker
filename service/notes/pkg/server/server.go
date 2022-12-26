package server

import (
	"context"
	"errors"
	"fmt"

	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
	"atypicaldev.com/conversation/notes/pkg/data"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type conversationServer struct {
	repository data.NotesRepository
	service_pb.UnimplementedConversationServiceServer
}

func newServer(repository data.NotesRepository) *conversationServer {
	return &conversationServer{
		repository: repository,
	}
}

func (s *conversationServer) GetNote(ctx context.Context, request *service_pb.GetNoteRequest) (response *service_pb.GetNoteResponse, err error) {
	note, err := s.repository.GetNote(request.NoteId)
	response = &service_pb.GetNoteResponse{Note: note}
	return
}

func (s *conversationServer) GetNotes(ctx context.Context, request *service_pb.GetNotesRequest) (response *service_pb.GetNotesResponse, err error) {
	notes, err := s.repository.GetNotes()
	l := ctxlogrus.Extract(ctx).Logger

	l.Printf("Notes from repository: %v", notes)

	response = &service_pb.GetNotesResponse{
		Notes: notes,
	}
	return
}

func (s *conversationServer) GetConversation(ctx context.Context, request *service_pb.GetConversationRequest) (response *service_pb.GetConversationResponse, err error) {
	convo, err := s.repository.GetConversation(request.GetConversationId())
	l := ctxlogrus.Extract(ctx).Logger
	l.Warnf("Fetching convo with id: [%s]", request.GetConversationId())
	if errors.Is(err, data.NewQueryError(data.FindOneErr)) {
		err = status.Errorf(codes.NotFound, err.Error())
	}
	response = &service_pb.GetConversationResponse{Conversation: convo}
	return
}

func (s *conversationServer) ListConversations(
	ctx context.Context,
	request *service_pb.ListConversationsRequest,
) (response *service_pb.ListConversationsResponse, err error) {
	l := ctxlogrus.Extract(ctx).Logger
	convos, err := s.repository.ListConversations()
	l.Infof("Repo Conversations: %v", convos)
	response = &service_pb.ListConversationsResponse{Conversations: convos}
	return
}

func (s *conversationServer) CreateNote(ctx context.Context, request *service_pb.CreateNoteRequest) (response *service_pb.CreateNoteResponse, err error) {
	fmt.Printf("Creating Note: %v\n", request)
	note, err := s.repository.CreateNote(request)
	response = &service_pb.CreateNoteResponse{Note: note}
	return
}

func (s *conversationServer) CreateConversation(
	ctx context.Context,
	request *service_pb.CreateConversationRequest,
) (response *service_pb.CreateConversationResponse, err error) {
	l := ctxlogrus.Extract(ctx).Logger
	l.Info("\n\nCreating new Conversation")
	convo, err := s.repository.CreateConversation(request)
	l.Infof("\n\nCreated new Conversation:\n%v\n\n", convo)
	response = &service_pb.CreateConversationResponse{Conversation: convo}
	return
}

func (s *conversationServer) UpdateReply(
	ctx context.Context,
	request *service_pb.UpdateReplyRequest,
) (response *service_pb.UpdateReplyResponse, err error) {
	l := ctxlogrus.Extract(ctx).Logger
	l.Info("\n\nUpdating reply")
	note, err := s.repository.UpdateReply(request)
	response = &service_pb.UpdateReplyResponse{Note: note}
	return
}
