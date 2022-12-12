'use client';

import { Note } from "../../../models/notes";
import { Card } from "../../../components/card";
import { NewNoteWrapper, NewNoteContext } from "./notes_context";
import { useContext } from "react";
import AddNoteButton from "./add_note_button";

type NotesViewProps = {
  notes: Note[],
}

export default function NotesView({ notes }: NotesViewProps) {
  const { enabled } = useContext(NewNoteContext);
  return <NewNoteWrapper>

    {notes.map(note => <Card key={note.id} content={note.content} reply={note.reply} />)}
    {enabled ? <div>New Note!</div> : null}
    <AddNoteButton />
  </NewNoteWrapper>
} 