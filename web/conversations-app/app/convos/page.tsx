class Note {
  constructor(
    public reply: string,
    public id: string,
    public content: string,
  ) { }
  static fromJson(): Note {
    return new Note('a', 'b', 'c');
  }

}

type NoteJson = { notes: Array<Note> }

async function getNotes(): Promise<NoteJson> {
  const res = await fetch('http://localhost:1337/notes');

  return JSON.parse(await res.json()) as Promise<NoteJson>;

}

export default async function Page() {
  const notes = (await getNotes())['notes'];
  // console.log(res['notes'])
  console.log(notes)
  return <div>
    <h1>Hello Conversation Page!</h1>
    <ul>
      {notes.map(note => <li key={note.id}>{note.content}</li>)}
    </ul>
  </div>
}