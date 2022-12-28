'use client';
import { useState } from "react";

type ReplySectionProps = {
  reply?: string | undefined
}

export default function ReplySection({ reply }: ReplySectionProps) {
  /// TODO(GikuyuNderitu): Mutate and update with new note data
  const [editing, setEditing] = useState(false)
  const editButtonText = editing ? 'Alter reply' : 'Edit reply'
  const addButtonText = editing ? 'Cancel' : 'Add reply'
  const hasReply = reply !== undefined && reply !== '';
  return (
    <div className="flex flex-col items-start">
      {hasReply ?
        <button className="text-blue-500" onClick={() => setEditing(!editing)}>{editButtonText}</button> :
        <button className="text-blue-500" onClick={() => setEditing(!editing)}>{addButtonText}</button>}
      {editing ?
        <textarea /> :
        <p className="text-base">{reply}</p>
      }
    </div>
  )
}