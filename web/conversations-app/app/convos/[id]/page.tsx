import { ConvoModel, ConvoModelJson } from "../../../models/convos";
import unmarshall, { Shaper } from "../../../util/unmarshal";
import AddNoteButton from "./add_note_button";
import NotesView from "./notes_view";

type ConvoDetailPageParams = {
  params: { id: string },
}

class ConversationShaper implements Shaper<ConvoModelJson> {
  shape(t: ConvoModelJson): ConvoModelJson {
    const model = t.conversation;

    return {
      conversation: new ConvoModel(model.title, model.id, model.notes),
    }
  }
}

async function getConvo(id: string): Promise<ConvoModel> {
  const res = await fetch(
    `http://localhost:1337/convos/${id}`,
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
      <NotesView notes={convo.notes} />
    </div>
  )
}