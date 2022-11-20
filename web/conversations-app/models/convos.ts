export class Convo {
  constructor(
    public title: string,
    public id: string,
  ) { }
}

export type ConvoJson = { conversations: Array<Convo> }