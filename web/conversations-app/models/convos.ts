export class Convo {
  constructor(
    public title: string,
    public id: string,
  ) { }
}

export type ConvosJson = { conversations: Array<Convo> }