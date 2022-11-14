class Note {
  constructor(
    public reply: string,
    public id: string,
    public content: string,
  ) { }
}

type NoteJson = { notes: Array<Note> }

async function getNotes(): Promise<Array<Note>> {
  const res = await fetch('http://localhost:1337/notes', { cache: 'no-store' },);

  return ((JSON.parse(await res.json())) as NoteJson)['notes'];

}

export default async function Page() {
  const notes = await getNotes();
  console.log(notes)
  return <div>
    <h1>Hello Conversation Page!</h1>
    <ul>
      {notes.map(note => <li key={note.id}>{note.content}</li>)}
    </ul>
  </div>
}