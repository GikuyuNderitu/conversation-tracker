package data

import (
	"errors"
	"log"

	pb "atypicaldev.com/conversation/api/notes"
	"github.com/surrealdb/surrealdb.go"
)

var (
	ErrInvalidResponse = errors.New("invalid SurrealDB response")
	ErrQuery           = errors.New("error occurred processing the SurrealDB query")
)

type NotesRepository interface {
	// Read operations
	GetNote(noteId string) (*pb.Note, error)
	GetNotes() ([]*pb.Note, error)
	ListConversations() ([]*pb.Conversation, error)
	GetConversation(convoId string) (*pb.Conversation, error)

	// Create operations
	CreateNote(request *pb.CreateNoteRequest) (*pb.Note, error)
	CreateConversation(request *pb.CreateConversationRequest) (*pb.Conversation, error)
}

type noteRepository struct {
	connectionUrl string
	dbEnvironment string
	// TODO(atypicaldev): Add logging
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

func (r noteRepository) GetNote(noteId string) (*pb.Note, error) {
	db := r.openConnection()
	defer db.Close()

	todoData, err := db.Query(todoQuery, map[string]interface{}{
		"id": noteId,
	})
	if err != nil {
		return nil, err
	}

	var note *pb.Note
	_, err = unmarshalRaw(todoData, &note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (r noteRepository) GetNotes() ([]*pb.Note, error) {
	db := r.openConnection()
	defer db.Close()

	todoData, err := db.Select(todoTable)
	if err != nil {
		return nil, err
	}

	var note []*pb.Note
	err = unmarshal(todoData, &note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (r noteRepository) GetConversation(convoId string) (*pb.Conversation, error) {
	db := r.openConnection()
	defer db.Close()

	todoData, err := db.Query(convoQuery, map[string]interface{}{
		"id": convoId,
	})
	if err != nil {
		return nil, err
	}

	var convo *pb.Conversation
	_, err = unmarshalRaw(todoData, &convo)
	if err != nil {
		return nil, err
	}
	return convo, nil
}

func (r noteRepository) ListConversations() ([]*pb.Conversation, error) {
	db := r.openConnection()
	defer db.Close()

	convoData, err := db.Select(convoTable)
	if err != nil {
		return nil, err
	}

	var conversations []*pb.Conversation
	err = unmarshal(convoData, &conversations)

	if err != nil {
		return nil, err
	}

	return conversations, nil
}

func (r noteRepository) CreateNote(request *pb.CreateNoteRequest) (*pb.Note, error) {
	// TODO(#5): Validate the request (convoId populated, content populated
	// non-empty, reply populated/empty string)
	db := r.openConnection()
	defer db.Close()

	noteData, err := db.Create(todoTable, map[string]any{
		"content":        request.Content,
		"reply":          request.Reply,
		"conversationId": request.ConversationId,
	})

	if err != nil {
		return nil, err
	}

	var note pb.Note
	err = unmarshal(noteData, &note)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r noteRepository) CreateConversation(request *pb.CreateConversationRequest) (*pb.Conversation, error) {
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

	var conversation pb.Conversation
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
