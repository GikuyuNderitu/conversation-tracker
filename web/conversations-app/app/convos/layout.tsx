export default function ConvoLayout({
  children
}: { children: React.ReactNode }) {

  return (
    <div>
      <nav className="h-16 bg-on-primary-surface flex items-center">
        <div className="mx-6">
          <h1 className="text-4xl text-on-primary-surface-text">Conversations</h1>
        </div>
      </nav>
      <main className="m-6">

        {children}
      </main>
    </div>
  )
}