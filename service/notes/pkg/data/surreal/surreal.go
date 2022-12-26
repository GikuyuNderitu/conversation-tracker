package surreal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	convo_pb "atypicaldev.com/conversation/notes/internal/proto/conversations/v1"
	notes_pb "atypicaldev.com/conversation/notes/internal/proto/notes/v1"
	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
	data_errors "atypicaldev.com/conversation/notes/pkg/data/errors"
	"github.com/surrealdb/surrealdb.go"
)

var (
	ErrInvalidResponse = errors.New("invalid SurrealDB response")
	ErrQuery           = errors.New("error occurred processing the SurrealDB query")
)

type surrealNoteRepository struct {
	connectionUrl string
	dbEnvironment string
}

func NewSurrealRepository(connectionUrl, dbEnv string) surrealNoteRepository {
	repository := surrealNoteRepository{
		connectionUrl: connectionUrl,
		dbEnvironment: dbEnv,
	}

	conn := repository.openConnection()
	conn.Close()

	return repository
}

func (r surrealNoteRepository) GetNote(noteId string) (*notes_pb.Note, error) {
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

func (r surrealNoteRepository) GetNotes() ([]*notes_pb.Note, error) {
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

func (r surrealNoteRepository) GetConversation(convoId string) (*convo_pb.Conversation, error) {
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
		return nil, data_errors.NewQueryError(data_errors.FindOneErr, convoTable)
	}
	return convo, nil
}

func (r surrealNoteRepository) ListConversations() ([]*convo_pb.Conversation, error) {
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

func (r surrealNoteRepository) CreateNote(request *service_pb.CreateNoteRequest) (*notes_pb.Note, error) {
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

func (r surrealNoteRepository) CreateConversation(request *service_pb.CreateConversationRequest) (*convo_pb.Conversation, error) {
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

func (r surrealNoteRepository) UpdateReply(request *service_pb.UpdateReplyRequest) (*notes_pb.Note, error) {
	panic("Unimplemented Error: Surreal repository UpdateReply not implemented")
}

func (r surrealNoteRepository) openConnection() *surrealdb.DB {
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

func (r surrealNoteRepository) signin(db *surrealdb.DB) bool {
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

// Unmarshal loads a SurrealDB response into a struct.
func unmarshal(data, v interface{}) error {
	var ok bool

	assertedData, ok := data.([]interface{})
	if !ok {
		return ErrInvalidResponse
	}
	sliceFlag := isSlice(v)

	var jsonBytes []byte
	var err error
	if !sliceFlag && len(assertedData) > 0 {
		jsonBytes, err = json.Marshal(assertedData[0])
	} else {
		jsonBytes, err = json.Marshal(assertedData)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, v)
	if err != nil {
		return err
	}

	return err
}

// UnmarshalRaw loads a raw SurrealQL response returned by Query into a struct. Queries that return with results will
// return ok = true, and queries that return with no results will return ok = false.
func unmarshalRaw(rawData, v interface{}) (ok bool, err error) {
	var data []interface{}
	if data, ok = rawData.([]interface{}); !ok {
		return false, ErrInvalidResponse
	}

	var responseObj map[string]interface{}
	if responseObj, ok = data[0].(map[string]interface{}); !ok {
		return false, ErrInvalidResponse
	}

	var status string
	if status, ok = responseObj["status"].(string); !ok {
		return false, ErrInvalidResponse
	}
	if status != statusOK {
		return false, ErrQuery
	}

	result := responseObj["result"]
	if len(result.([]interface{})) == 0 {
		return false, nil
	}
	err = unmarshal(result, v)
	if err != nil {
		return false, err
	}

	return true, nil
}

func isSlice(possibleSlice interface{}) bool {
	slice := false

	switch v := possibleSlice.(type) { //nolint:gocritic
	default:
		res := fmt.Sprintf("%s", v)
		if res == "[]" || res == "&[]" || res == "*[]" {
			slice = true
		}
	}

	return slice
}
