import AddConvoButton from "./add_convo_button";
import { ConvoModel } from "../../models/convos";

export default function LoadedState({ convos }: { convos: Array<ConvoModel> }) {
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