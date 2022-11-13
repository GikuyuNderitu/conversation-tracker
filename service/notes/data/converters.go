package data

import pb "atypicaldev.com/conversation/api/notes"

type note struct {
	id, convoId, content, reply, parent string
	// TODO(atypicaldev): Handle parse of children
	children []*note // nolint
}
type convo struct {
	id, title string
	notes     []*note
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
