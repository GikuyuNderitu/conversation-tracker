export default function ConvoLayout({
  children
}: { children: React.ReactNode }) {

  return (
    <div className="min-h-screen flex flex-col">
      <nav className="h-16 bg-on-surface flex items-center">
        <div className="mx-6">
          <h1 className="text-4xl text-on-surface-text">Conversations</h1>
        </div>
      </nav>
      <main className="mx-6 min-h-full flex grow">
        {children}
      </main>
    </div>
  )
}