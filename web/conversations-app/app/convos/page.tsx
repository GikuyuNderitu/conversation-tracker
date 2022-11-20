import unmarshall from "../../util/unmarshal";
import AddConvoButton from "./add_convo_button";

class Note {
  constructor(
    public reply: string,
    public id: string,
    public content: string,
  ) { }
}

type NoteJson = { notes: Array<Note> }

async function getNotes(): Promise<Array<Note>> {
  const res = await fetch(
    'http://localhost:1337/notes',
    { cache: 'no-store' },
  );

  return (await unmarshall<NoteJson>(res))['notes'];
}

export default async function Page() {
  const notes = await getNotes();
  console.log(notes)


  return notes.length > 0 ? <EmptyState /> : <LoadedState notes={notes} />;
}

function EmptyState() {
  return (
    <div className="flex justify-center items-center min-h-full">
      <h2 className="text-3xl text-primary">Add a conversation to get started!</h2>
      <AddConvoButton />
    </div>
  )
}

function LoadedState({ notes }: { notes: Array<Note> }) {
  return (
    <div>
      <h1>Hello Conversation Page!</h1>
      <ul>
        {notes.map(note => <li key={note.id}>{note.content}</li>)}
      </ul>

      <AddConvoButton />
    </div>
  );
}