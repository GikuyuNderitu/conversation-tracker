package postgres

import (
	"strconv"

	convo_pb "atypicaldev.com/conversation/notes/internal/proto/conversations/v1"
	notes_pb "atypicaldev.com/conversation/notes/internal/proto/notes/v1"
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Title string
	Notes []Note
}

type Note struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	Conversation   Conversation
	ConversationID uint
	Content        string `gorm:"not null"`
	Reply          string
	Note           []Note `gorm:"foreignKey:Parent"`
	Parent         *int
}

func (c Conversation) toPb() *convo_pb.Conversation {
	convoId := strconv.Itoa(int(c.ID))
	return &convo_pb.Conversation{
		Id:    convoId,
		Title: c.Title,
		Notes: convertNotestoPb(c.Notes),
	}
}

func (n Note) toPb() *notes_pb.Note {
	// parent := n.Parent
	return &notes_pb.Note{
		Content: n.Content,
		Reply:   n.Reply,
		// Parent:         strconv.Itoa(parent),
		Id:             strconv.Itoa(int(n.ID)),
		ConversationId: strconv.Itoa(int(n.ConversationID)),
	}
}

func convertNotestoPb(notes []Note) []*notes_pb.Note {
	pbs := make([]*notes_pb.Note, 0, len(notes))
	for _, note := range notes {
		pbs = append(pbs, note.toPb())
	}
	return pbs
}
