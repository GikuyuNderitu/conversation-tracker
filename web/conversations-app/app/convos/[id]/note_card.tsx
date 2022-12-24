import Card from "../../../components/card";
import ReplySection from "./reply_section";

type NoteCardProps = {
  content: string | undefined;
  reply?: string | undefined;
}

export default function NoteCard({ content, reply }: NoteCardProps) {
  return (
    <Card className="my-8">
      <p className="text-xl">{content}</p>

      <ReplySection reply={reply} />
    </Card>
  )
}