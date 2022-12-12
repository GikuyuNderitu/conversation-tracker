'use client';

import { Note } from "../../../models/notes";
import { NewNoteWrapper } from "./notes_context";
import AddNoteButton from "./add_note_button";
import NoteCard from "./note_card";
import NewNote from "./new_note";

type NotesViewProps = {
  notes: Note[],
}

export default function NotesView({ notes }: NotesViewProps) {
  return (
    <NewNoteWrapper>
      {notes.map(note => <NoteCard key={note.id} content={note.content} reply={note.reply} />)}
      <NewNote />
      <AddNoteButton />
    </NewNoteWrapper>
  )
} 