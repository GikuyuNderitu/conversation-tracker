import { ConvoModel, ConvoModelJson } from "../../../models/convos";
import unmarshall from "../../../util/unmarshal";
import NotesView from "./notes_view";
import ConversationShaper from './conversation_shaper';

type ConvoDetailPageParams = {
  params: { id: string },
}


async function getConvo(id: string): Promise<ConvoModel> {
  const res = await fetch(
    `http://localhost:1337/v1/conversations/${id}`,
    { cache: 'no-store' },
  );

  const unmarshalled = await unmarshall<ConvoModelJson>(res, new ConversationShaper())
  console.log(unmarshalled)
  return unmarshalled['conversation'];
}

export default async function Page({ params: { id } }: ConvoDetailPageParams) {
  const convo = await getConvo(id);
  console.log(convo)

  return (
    <div>
      <h1 className="text-2xl">
        {convo.title}
      </h1>
      <NotesView convo={convo} />
    </div>
  )
}