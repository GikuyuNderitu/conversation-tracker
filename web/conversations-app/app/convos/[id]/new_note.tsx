'use client'

import { useMutation } from "@tanstack/react-query";
import { useAtom } from "jotai";
import { usePathname } from "next/navigation";
import React, { useRef, useState } from "react"
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

type NewNoteProps = {
  reload: VoidFunction;
}

export default function NewNote({ reload }: NewNoteProps) {
  const [enabled] = useAtom(newNoteAtom);
  const [noteContent, setNoteContent] = useState('');
  const textAreaRef = useRef<HTMLTextAreaElement>(null)
  const { mutate } = useMutation({
    mutationFn: createNote,
    onSuccess() {
      setNoteContent('')
      textAreaRef.current?.focus()
      reload();
    },
  });

  const pathname = usePathname();

  const conversationId = pathname?.split('/')[2] || '';

  const updateTextArea = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setNoteContent(e.target.value);
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.ctrlKey && e.key === "Enter") {
      console.log("Trying to submit something")
      e.preventDefault()
      submit();
    }
  }

  const submit = () => {
    mutate({ content: noteContent, conversationId })
  }

  return (
    enabled ? <Card className="bg-white flex flex-col">
      <textarea ref={textAreaRef} onKeyDown={handleKeyDown} autoFocus className="form-textarea" value={noteContent} onChange={updateTextArea} />
      <button className="self-end mt-2 bg-accent px-3 py-1 rounded" onClick={submit}>Submit</button>
    </Card> : null
  )
}