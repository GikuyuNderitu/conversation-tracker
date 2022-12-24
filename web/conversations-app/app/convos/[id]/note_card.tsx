'use client'
import Card from "../../../components/card";
import ReplySection from "./reply_section";
import { XCircleIcon } from '@heroicons/react/20/solid'

type NoteCardProps = {
  content: string | undefined;
  reply?: string | undefined;
}

export default function NoteCard({ content, reply }: NoteCardProps) {
  return (
    <Card className="my-8">
      {/* TODO(GikuyuNderitu): Add mutation to delete note */}
      <CloseButton onClick={() => console.log('closing')} />
      <p className="text-xl">{content}</p>

      <ReplySection reply={reply} />
    </Card>
  )
}

type CloseButtonProps = {
  onClick(event: React.MouseEvent<HTMLButtonElement>): void;
}

function CloseButton({ onClick }: CloseButtonProps) {
  return (
    <button className="
    group 
    float-right
    mr-2 ml-auto
      bg-stone-50 
      rounded-[50%]"
      onClick={onClick}>
      <XCircleIcon className="
      w-6 h-6
      text-blue-500 stroke-blue-500 stroke-0 group-hover:stroke-1
      transition-color duration-250" />
    </button>
  )
}