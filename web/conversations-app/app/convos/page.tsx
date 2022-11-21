import { ConvoModelsJson, ConvoModel } from '../../models/convos'
import unmarshall from "../../util/unmarshal";
import EmptyState from './_empty_page';
import LoadedState from "./_loaded_page";

async function getConvos(): Promise<Array<ConvoModel>> {
  const res = await fetch(
    'http://localhost:1337/convos',
    { cache: 'no-store' },
  );

  return (await unmarshall<ConvoModelsJson>(res))['conversations'];
}

export default async function Page() {
  const convos = await getConvos();
  console.log(convos)

  return convos.length == 0 ? <EmptyState /> : <LoadedState convos={convos} />;
}