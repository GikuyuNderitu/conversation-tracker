import { ConvoModelJson } from "../../../models/convos";
import { Shaper } from "../../../util/unmarshal";

export default class ConversationShaper implements Shaper<ConvoModelJson> {
  shape(t: ConvoModelJson): ConvoModelJson {
    const model = t.conversation;

    return {
      conversation: {
        title: model.title,
        id: model.id,
        notes: model.notes ?? [],
      }
    }
  }
}