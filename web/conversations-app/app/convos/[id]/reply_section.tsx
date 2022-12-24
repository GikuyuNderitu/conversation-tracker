'use client';
import { useState } from "react";

type ReplySectionProps = {
  reply?: string | undefined
}

export default function ReplySection({ reply }: ReplySectionProps) {
  /// TODO(GikuyuNderitu): Mutate and update with new note data
  const [editing, setEditing] = useState(false)
  const editButtonText = editing ? 'Alter reply' : 'Edit reply'
  const hasReply = reply !== undefined && reply !== '';
  return (
    <div className="flex flex-col items-start">
      {editing ?
        <textarea /> :
        <p className="text-base">{reply}</p>
      }
      {hasReply ?
        <button className="text-blue-500" onClick={() => setEditing(!editing)}>{editButtonText}</button> :
        <button className="text-blue-500" onClick={() => setEditing(!editing)}>Add reply</button>}
    </div>
  )
}