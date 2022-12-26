package data

import (
	convo_pb "atypicaldev.com/conversation/notes/internal/proto/conversations/v1"
	notes_pb "atypicaldev.com/conversation/notes/internal/proto/notes/v1"
	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
)

type NotesRepository interface {
	// Read operations
	GetNote(noteId string) (*notes_pb.Note, error)
	GetNotes() ([]*notes_pb.Note, error)
	ListConversations() ([]*convo_pb.Conversation, error)
	GetConversation(convoId string) (*convo_pb.Conversation, error)

	// Create operations
	CreateNote(request *service_pb.CreateNoteRequest) (*notes_pb.Note, error)
	CreateConversation(request *service_pb.CreateConversationRequest) (*convo_pb.Conversation, error)
	UpdateReply(request *service_pb.UpdateReplyRequest) (*notes_pb.Note, error)
}
