import { Note } from "./notes";

export class ConvoModel {
  constructor(
    public title: string,
    public id: string,
    public notes: Array<Note> = [],
  ) { }
}

export type ConvoModelsJson = { conversations: Array<ConvoModel> }
export type ConvoModelJson = { conversation: ConvoModel }