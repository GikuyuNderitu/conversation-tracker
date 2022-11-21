export class ConvoModel {
  constructor(
    public title: string,
    public id: string,
  ) { }
}

export type ConvoModelsJson = { conversations: Array<ConvoModel> }