package postgres

import (
	"log"

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

	var convos Conversation

	db.Find(&convos)
	return nil, nil
}

func (r psqlRepository) GetConversation(convoId string) (*convo_pb.Conversation, error) {
	return nil, nil
}

func (r psqlRepository) CreateNote(request *service_pb.CreateNoteRequest) (*notes_pb.Note, error) {
	return nil, nil

}

func (r psqlRepository) CreateConversation(request *service_pb.CreateConversationRequest) (*convo_pb.Conversation, error) {
	return nil, nil
}
