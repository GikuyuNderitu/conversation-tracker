package postgres

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	Title string
}

type Note struct {
	gorm.Model
	Conversation   Conversation
	ConversationID int
	Content        string `gorm:"not null"`
	Reply          string
	Note           []Note `gorm:"foreignKey:Parent"`
	Parent         int
}
