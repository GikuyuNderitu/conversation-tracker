export class Note {
  constructor(
    public reply: string,
    public id: string,
    public content: string,
  ) { }
}

export type NotesJson = { notes: Array<Note> }