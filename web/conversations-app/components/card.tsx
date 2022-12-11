type CardProps = {
  content?: string | undefined;
  reply?: string | undefined;
}

export function Card({ content, reply }: CardProps) {
  return (
    <div className="rounded">
      {content !== undefined ?
        <p className="text-xl">{content}</p> : null
      }
      {reply !== undefined ? <p className="text-base">{reply}</p> : null}
    </div>
  )
}