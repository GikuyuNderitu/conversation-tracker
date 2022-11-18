export default function ConvoLayout({
  children
}: { children: React.ReactNode }) {

  return (
    <div className="p-6">
      <nav></nav>
      <main>

        {children}
      </main>
    </div>
  )
}