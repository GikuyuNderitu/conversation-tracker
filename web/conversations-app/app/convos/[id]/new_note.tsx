'use client'

import { useAtom } from "jotai";
import { useState } from "react"
import Card from "../../../components/card";
import { newNoteAtom } from "./new_note_state";


export default function NewNote() {
  const [enabled] = useAtom(newNoteAtom);
  const [note, setNote] = useState('');

  const updateTextArea = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    console.log(e.target.value);
    setNote(e.target.value);
  }

  const submit = () => {
    // TODO(GikuyuNderitu): Launch mutation
    console.log("submitting new note")
  }


  return (
    enabled ? <Card className="bg-white flex flex-col">
      <textarea autoFocus className="form-textarea" value={note} onChange={updateTextArea} />
      <button className="self-end mt-2 bg-accent px-3 py-1 rounded" onClick={submit}>Submit</button>
    </Card> : null
  )
}