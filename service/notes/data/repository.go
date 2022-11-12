package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	pb "atypicaldev.com/conversation/notes/api"
	"github.com/surrealdb/surrealdb.go"
)

var (
	ErrInvalidResponse = errors.New("invalid SurrealDB response")
	ErrQuery           = errors.New("error occurred processing the SurrealDB query")
)

type note struct {
	id, convoId, content, reply, parent string
	children                            []*note // nolint
}
type convo struct {
	id, title string
	notes     []*note
}

const (
	convoTable           = "convos"
	noteTable            = "notes"
	todoTable            = "todos"
	conversationDBPrefix = "conversations"
	conversationNSPrefix = conversationDBPrefix
	statusOK             = "OK"
)

var (
	todoQuery = fmt.Sprintf("SELECT * FROM %s WHERE id = $id", todoTable)
)

type NotesRepository interface {
	GetNote(noteId string) *pb.Note
	GetNotes(noteId string) []*pb.Note
	ListConversations() []pb.Conversation
	GetConversation(convoId string) *pb.Conversation
	CreateNote(request *pb.CreateNoteRequest) *pb.Note
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

func (r noteRepository) GetNote(noteId string) *pb.Note {
	db := r.openConnection()
	defer db.Close()

	todoData, err := db.Query(todoQuery, map[string]interface{}{
		"id": noteId,
	})
	if err != nil {
		return nil
	}

	var note *pb.Note
	_, err = unmarshalRaw(todoData, &note)
	if err != nil {
		return nil
	}
	return note
}

func (r noteRepository) GetNotes(noteId string) []*pb.Note {
	db := r.openConnection()
	defer db.Close()

	// todoData, err := db.Select(todoTable)
	todoData, err := db.Query(todoQuery, map[string]interface{}{
		"id": noteId,
	})
	if err != nil {
		return nil
	}

	var note []*pb.Note
	_, err = unmarshalRaw(todoData, &note)
	if err != nil {
		return nil
	}
	return note
}

func (r noteRepository) GetConversation(convoId string) *pb.Conversation {
	db := r.openConnection()
	defer db.Close()

	return &pb.Conversation{}
}

func (r noteRepository) ListConversations() []pb.Conversation {
	db := r.openConnection()
	defer db.Close()

	convoData, err := db.Select(convoTable)
	if err != nil {
		return nil
	}

	var conversations []convo
	err = unmarshal(convoData, &conversations)

	if err != nil {
		return nil
	}

	return convertDbConversationListToConvo(conversations)
}

func (r noteRepository) CreateNote(request *pb.CreateNoteRequest) *pb.Note {
	//TODO: Validate the request (convoId populated, content populated non-empty, reply populated/empty string)
	db := r.openConnection()
	defer db.Close()

	noteData, err := db.Create(todoTable, map[string]any{
		"content":        request.Content,
		"reply":          request.Reply,
		"conversationId": request.ConversationId,
	})

	if err != nil {
		return nil
	}

	var note pb.Note
	err = unmarshal(noteData, &note)
	if err != nil {
		return nil
	}

	return &note
}

func convertDbNotesListToProto(dbNotes []*note) []*pb.Note {
	var notes []*pb.Note
	for _, note := range dbNotes {
		notes = append(notes, convertDbNoteToProto(note))
	}

	return notes
}

func convertDbNoteToProto(dbNote *note) *pb.Note {
	if dbNote == nil {
		return nil
	}
	return &pb.Note{
		Id:             dbNote.id,
		ConversationId: dbNote.convoId,
		Content:        dbNote.content,
		Reply:          dbNote.reply,
		Parent:         dbNote.parent,
	}
}

func convertDbConversationListToConvo(dbConvos []convo) []pb.Conversation {
	var convos []pb.Conversation
	for _, convo := range dbConvos {
		convos = append(convos, convertDbConversationToConvo(convo))
	}

	return convos
}

func convertDbConversationToConvo(dbConvo convo) pb.Conversation {
	return pb.Conversation{
		Id:    dbConvo.id,
		Title: dbConvo.title,
		Notes: convertDbNotesListToProto(dbConvo.notes),
	}
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
