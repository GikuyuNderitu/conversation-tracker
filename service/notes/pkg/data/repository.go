package data

import (
	"errors"

	convo_pb "atypicaldev.com/conversation/notes/internal/proto/conversations/v1"
	notes_pb "atypicaldev.com/conversation/notes/internal/proto/notes/v1"
	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
)

var (
	ErrInvalidResponse = errors.New("invalid SurrealDB response")
	ErrQuery           = errors.New("error occurred processing the SurrealDB query")
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
}
