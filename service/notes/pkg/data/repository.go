package data

import (
	"errors"
	"log"

	convo_pb "atypicaldev.com/conversation/notes/internal/proto/conversations/v1"
	notes_pb "atypicaldev.com/conversation/notes/internal/proto/notes/v1"
	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
	"github.com/surrealdb/surrealdb.go"
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

type noteRepository struct {
	connectionUrl string
	dbEnvironment string
}

func NewRepository(connectionUrl, dbEnv string) noteRepository {
	repository := noteRepository{
		connectionUrl: connectionUrl,
		dbEnvironment: dbEnv,
	}

	conn := repository.openConnection()
	conn.Close()

	return repository
}

func (r noteRepository) GetNote(noteId string) (*notes_pb.Note, error) {
	db := r.openConnection()
	defer db.Close()

	todoData, err := db.Query(noteQuery, map[string]interface{}{
		"id": noteId,
	})
	if err != nil {
		return nil, err
	}

	var note *notes_pb.Note
	_, err = unmarshalRaw(todoData, &note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (r noteRepository) GetNotes() ([]*notes_pb.Note, error) {
	db := r.openConnection()
	defer db.Close()

	todoData, err := db.Select(todoTable)
	if err != nil {
		return nil, err
	}

	var note []*notes_pb.Note
	err = unmarshal(todoData, &note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (r noteRepository) GetConversation(convoId string) (*convo_pb.Conversation, error) {
	db := r.openConnection()
	defer db.Close()

	todoData, err := db.Query(convoQuery, map[string]interface{}{
		"id": convoId,
	})
	if err != nil {
		return nil, err
	}

	var convo *convo_pb.Conversation
	_, err = unmarshalRaw(todoData, &convo)
	if err != nil {
		return nil, err
	}

	if convo == nil {
		return nil, NewQueryError(FindOneErr, convoTable)
	}
	return convo, nil
}

func (r noteRepository) ListConversations() ([]*convo_pb.Conversation, error) {
	db := r.openConnection()
	defer db.Close()

	convoData, err := db.Select(convoTable)
	if err != nil {
		return nil, err
	}

	var conversations []*convo_pb.Conversation
	err = unmarshal(convoData, &conversations)

	if err != nil {
		return nil, err
	}

	return conversations, nil
}

func (r noteRepository) CreateNote(request *service_pb.CreateNoteRequest) (*notes_pb.Note, error) {
	// TODO(#5): Validate the request (convoId populated, content populated
	// non-empty, reply populated/empty string)
	db := r.openConnection()
	defer db.Close()

	noteData, err := db.Query(createNoteQuery, map[string]any{
		"content": request.Content,
		"reply":   request.Reply,
		"convo":   request.ConversationId,
	})

	if err != nil {
		return nil, err
	}

	var note notes_pb.Note
	err = unmarshal(noteData, &note)
	if err != nil {
		return nil, err
	}

	_, err = db.Query(relateNoteWithConvo, map[string]any{
		"note":  note.GetId(),
		"convo": request.GetConversationId(),
	})

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r noteRepository) CreateConversation(request *service_pb.CreateConversationRequest) (*convo_pb.Conversation, error) {
	db := r.openConnection()
	defer db.Close()

	if err := validateConvoRequest(request); err != nil {
		return nil, err
	}

	noteData, err := db.Create(convoTable, map[string]any{
		"title": request.Title,
	})

	if err != nil {
		return nil, err
	}

	var conversation convo_pb.Conversation
	err = unmarshal(noteData, &conversation)
	if err != nil {
		return nil, err
	}

	return &conversation, nil
}

func (r noteRepository) openConnection() *surrealdb.DB {
	db, err := surrealdb.New(r.connectionUrl)

	if err != nil {
		log.Fatalf("Problem establishing connection to Database, %v", err)
	}

	if ok := r.signin(db); !ok {
		return nil
	}

	_, err = db.Use(conversationNSPrefix+r.dbEnvironment, conversationDBPrefix+r.dbEnvironment)

	if err != nil {
		return nil
	}

	return db
}

func (r noteRepository) signin(db *surrealdb.DB) bool {
	_, err := db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	})

	if err != nil {
		log.Fatal("Error Signing into database")
		return false
	}

	return true
}
