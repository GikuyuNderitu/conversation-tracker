'use client';

import { ConvoModel } from "../../models/convos";

type ConvoProps = {
  convo: ConvoModel
}

export default function Convo({
  convo
}: ConvoProps) {
  return (
    <div className="w-[280px] h-[130px] p-4 bg-on-surface text-on-surface-text rounded-3xl">
      <h4 className="font-medium text-xl">{convo.title}</h4>
    </div>
  )
}