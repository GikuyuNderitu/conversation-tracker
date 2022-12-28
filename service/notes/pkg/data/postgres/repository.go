package postgres

import (
	"log"
	"strconv"

	convo_pb "atypicaldev.com/conversation/notes/internal/proto/conversations/v1"
	notes_pb "atypicaldev.com/conversation/notes/internal/proto/notes/v1"
	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type psqlRepository struct {
	connectionUrl string
}

func MigratePostgres(connectionUrl string) {
	db, err := psqlRepository{connectionUrl: connectionUrl}.oppenConnection()
	if err != nil {
		log.Fatalf("Encountered an error while openning connection: %v", err)
	}

	db.AutoMigrate(&Conversation{}, &Note{})
}

func NewPsqlRepository(connectionurl string) *psqlRepository {
	return &psqlRepository{
		connectionUrl: connectionurl,
	}
}

func (r psqlRepository) oppenConnection() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(r.connectionUrl), &gorm.Config{})
}

// Read operations
func (r psqlRepository) GetNote(noteId string) (*notes_pb.Note, error) {
	return nil, nil
}

func (r psqlRepository) GetNotes() ([]*notes_pb.Note, error) {
	return nil, nil
}

func (r psqlRepository) ListConversations() ([]*convo_pb.Conversation, error) {
	db, err := r.oppenConnection()
	if err != nil {
		return nil, err
	}

	var convos []Conversation

	db.Find(&convos)

	conversations := make([]*convo_pb.Conversation, 0, len(convos))
	for _, convo := range convos {
		conversations = append(conversations, &convo_pb.Conversation{
			Id:    strconv.Itoa(int(convo.ID)),
			Title: convo.Title,
		})
	}
	return conversations, nil
}

func (r psqlRepository) GetConversation(convoId string) (*convo_pb.Conversation, error) {
	id, err := strconv.Atoi(convoId)
	if err != nil {
		return nil, err
	}

	db, err := r.oppenConnection()
	if err != nil {
		return nil, err
	}

	convo := &Conversation{ID: uint(id)}
	res := db.Model(&Conversation{}).Preload("Notes").Find(convo)
	if res.Error != nil {
		return nil, res.Error
	}
	return &convo_pb.Conversation{
		Id:    convoId,
		Title: convo.Title,
		Notes: convertNotestoPb(convo.Notes),
	}, nil
}

func (r psqlRepository) CreateNote(request *service_pb.CreateNoteRequest) (*notes_pb.Note, error) {
	convoId, err := strconv.Atoi(request.ConversationId)
	if err != nil {
		return nil, err
	}
	db, err := r.oppenConnection()
	if err != nil {
		return nil, err
	}

	note := Note{
		Content:        request.Content,
		Reply:          request.Reply,
		ConversationID: uint(convoId),
	}

	res := db.Create(&note)
	if res.Error != nil {
		return nil, res.Error
	}

	err = db.Model(&Conversation{
		ID: uint(convoId),
	}).Association("Notes").Append(&note)

	if err != nil {
		return nil, err
	}

	return &notes_pb.Note{
		Content:        note.Content,
		Reply:          note.Reply,
		ConversationId: request.ConversationId,
	}, nil
}

func (r psqlRepository) CreateConversation(request *service_pb.CreateConversationRequest) (*convo_pb.Conversation, error) {
	db, err := r.oppenConnection()
	if err != nil {
		return nil, err
	}

	convo := Conversation{
		Title: request.Title,
	}

	res := db.Create(&convo)
	if res.Error != nil {
		return nil, res.Error
	}

	return &convo_pb.Conversation{
		Title: convo.Title,
		Id:    strconv.Itoa(int(convo.ID)),
	}, nil
}

func (r psqlRepository) UpdateReply(request *service_pb.UpdateReplyRequest) (*notes_pb.Note, error) {
	noteId, err := getDBID(request.NoteId)
	if err != nil {
		return nil, err
	}

	db, err := r.oppenConnection()
	if err != nil {
		return nil, err
	}

	note := &Note{ID: noteId}

	res := db.First(note).Update("Reply", request.Reply)
	if res.Error != nil {
		return nil, res.Error
	}

	return note.toPb(), nil
}

func (r psqlRepository) DeleteNote(request *service_pb.DeleteNoteRequest) (*convo_pb.Conversation, error) {
	noteId, err := getDBID(request.NoteId)
	if err != nil {
		return nil, err
	}

	convoId, err := getDBID(request.ConversationId)
	if err != nil {
		return nil, err
	}

	db, err := r.oppenConnection()
	if err != nil {
		return nil, err
	}

	dbConvo := &Conversation{ID: convoId}
	dbNote := &Note{ID: noteId}

	txn := db.Delete(dbNote)
	if txn.Error != nil {
		return nil, txn.Error
	}

	txn = db.Find(dbConvo)
	if txn.Error != nil {
		return nil, txn.Error
	}

	return dbConvo.toPb(), nil
}

func getDBID(requestId string) (uint, error) {
	id, err := strconv.Atoi(requestId)
	if err != nil {
		log.Printf("Error converting ID in request: Id=%s", requestId)
		return 0, err
	}
	return uint(id), nil
}
