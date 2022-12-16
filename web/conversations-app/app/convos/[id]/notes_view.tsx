'use client';

import AddNoteButton from "./add_note_button";
import NoteCard from "./note_card";
import NewNote from "./new_note";
import { Note } from "../../../models/notes";
import { useQuery } from "@tanstack/react-query";
import { ConvoModel, ConvoModelJson } from "../../../models/convos";
import unmarshall from "../../../util/unmarshal";
import ConversationShaper from "./conversation_shaper";

type NotesViewProps = {
  convo: ConvoModel,
}

async function getConvo(convo: ConvoModel): Promise<ConvoModel> {
  const res = await fetch(`http://localhost:1337/convos/${convo.id}`, {
    method: 'GET',
    headers: {
      'Content-type': 'application/json',
    }
  });

  return (await unmarshall<ConvoModelJson>(res, new ConversationShaper())).conversation
}

async function getNotes(convo: ConvoModel): Promise<Note[]> {
  if (convo.id === undefined) return [];
  const convoModel = await getConvo(convo);

  return convoModel.notes;
}

export default function NotesView({ convo }: NotesViewProps) {
  const { data } = useQuery({
    queryKey: ['convo'],
    queryFn: () => getNotes(convo),
    initialData: convo.notes,
  });
  return (
    <>
      {data.map(note => <NoteCard key={note.id} content={note.content} reply={note.reply} />)}
      <NewNote />
      <AddNoteButton />
    </>
  )
} 