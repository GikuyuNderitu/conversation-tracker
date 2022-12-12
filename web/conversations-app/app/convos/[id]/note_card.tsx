import Card from "../../../components/card";

type NoteCardProps = {
  content?: string | undefined;
  reply?: string | undefined;
}

export default function NoteCard({ content, reply }: NoteCardProps) {
  return <Card>
    {content !== undefined ?
      <p className="text-xl">{content}</p> : null
    }
    {reply !== undefined ? <p className="text-base">{reply}</p> : null}
  </Card>
}