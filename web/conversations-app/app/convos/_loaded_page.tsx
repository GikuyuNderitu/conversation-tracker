import AddConvoButton from "./add_convo_button";
import { ConvoModel } from "../../models/convos";
import Convo from "./convo";

export default function LoadedState({ convos }: { convos: Array<ConvoModel> }) {
  console.log(convos)
  // TODO (GikuyuNderitu): Fix grid layout
  return (
    <div className="w-full grid grid-flow-col shrink-0 justify-start grid-rows-6 pt-4 gap-9">
      {
        convos.map(convo => <Convo key={convo.id} convo={convo} />)
      }
      <AddConvoButton />
    </div>
  );
}