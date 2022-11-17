import { PlusIcon } from '@heroicons/react/24/solid'
export default function ConvoLayout({
  children
}: { children: React.ReactNode }) {

  return (
    <main>
      <div>
        {children}
      </div>
      <button className="fixed bottom-0 right-0 mb-16 mr-16 bg-stone-50 rounded-lg">
        <PlusIcon className="w-20 h-20 text-blue-500" />
      </button>
    </main>
  )
}