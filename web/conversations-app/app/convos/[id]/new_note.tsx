'use client'

import { useMutation } from "@tanstack/react-query";
import { useAtom } from "jotai";
import { usePathname } from "next/navigation";
import { useState } from "react"
import Card from "../../../components/card";
import { Note } from '../../../models/notes';
import { newNoteAtom } from "./new_note_state";

type NewNote = {
  conversationId: string,
  parentNoteId?: string,
  content: string,
  reply?: string,
}

async function createNote(newNote: NewNote): Promise<Note> {
  const options = {
    method: "POST",
    body: JSON.stringify(newNote),
    headers: {
      'Content-type': 'application/json',
    },
  }

  const data = await (await fetch(`http://localhost:1337/v1/conversations/${newNote.conversationId}/notes`, options)).json()
  return data as Note;
}


export default function NewNote() {
  const [enabled] = useAtom(newNoteAtom);
  const [noteContent, setNoteContent] = useState('');
  const { mutate } = useMutation({
    mutationFn: createNote,
  });

  const pathname = usePathname();

  const conversationId = pathname?.split('/')[2] || '';

  console.log(`Pathname: ${pathname}`)
  console.log(`ConversationId: ${conversationId}`)

  const updateTextArea = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    console.log(e.target.value);
    setNoteContent(e.target.value);
  }

  const submit = () => {
    // TODO(GikuyuNderitu): Launch mutation
    console.log("submitting new note")
    mutate({ content: noteContent, conversationId })
  }


  return (
    enabled ? <Card className="bg-white flex flex-col">
      <textarea autoFocus className="form-textarea" value={noteContent} onChange={updateTextArea} />
      <button className="self-end mt-2 bg-accent px-3 py-1 rounded" onClick={submit}>Submit</button>
    </Card> : null
  )
}