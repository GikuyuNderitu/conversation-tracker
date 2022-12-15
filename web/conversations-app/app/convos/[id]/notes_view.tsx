'use client';

import { useState } from "react";

import { NewNoteWrapper } from "./notes_context";
import AddNoteButton from "./add_note_button";
import NoteCard from "./note_card";
import NewNote from "./new_note";
import { Note } from "../../../models/notes";
import { useQuery } from "@tanstack/react-query";
import { ConvoModel } from "../../../models/convos";

type NotesViewProps = {
  convo: ConvoModel,
}

async function getNotes(convo: ConvoModel): Promise<Note[]> {
  return [];
}

export default function NotesView({ convo }: NotesViewProps) {
  const { data } = useQuery({
    queryKey: ['convo'],
    queryFn: () => getNotes(convo),
    initialData: convo.notes,
  });
  return (
    <NewNoteWrapper>
      {data.map(note => <NoteCard key={note.id} content={note.content} reply={note.reply} />)}
      <NewNote />
      <AddNoteButton />
    </NewNoteWrapper>
  )
} 