import AddConvoButton from "./add_convo_button";

export default function EmptyState() {
  return (
    <div className="flex justify-center items-center min-h-full">
      <h2 className="text-3xl text-primary">Add a conversation to get started!</h2>
      <AddConvoButton />
    </div>
  )
}