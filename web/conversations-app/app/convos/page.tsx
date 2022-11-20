import unmarshall from "../../util/unmarshal";
import AddConvoButton from "./add_convo_button";
import { ConvoJson, Convo } from '../../models/convos'

async function getConvos(): Promise<Array<Convo>> {
  const res = await fetch(
    'http://localhost:1337/convos',
    { cache: 'no-store' },
  );

  return (await unmarshall<ConvoJson>(res))['conversations'];
}

export default async function Page() {
  const convos = await getConvos();
  console.log(convos)


  return convos.length == 0 ? <EmptyState /> : <LoadedState convos={convos} />;
}

function EmptyState() {
  return (
    <div className="flex justify-center items-center min-h-full">
      <h2 className="text-3xl text-primary">Add a conversation to get started!</h2>
      <AddConvoButton />
    </div>
  )
}

function LoadedState({ convos }: { convos: Array<Convo> }) {
  return (
    <div>
      <h1>Hello Conversation Page!</h1>
      <ul>
        {convos.map(convo => <li key={convo.id}>{convo.title}</li>)}
      </ul>

      <AddConvoButton />
    </div>
  );
}