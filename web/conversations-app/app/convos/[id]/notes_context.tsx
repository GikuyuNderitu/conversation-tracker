
import { createContext, useState } from "react";

export const NewNoteContext = createContext({ enabled: false, toggle: () => { } });

export function NewNoteWrapper({ children }: { children: React.ReactNode }) {
  const [enabled, setEnabled] = useState(false);
  const toggle = () => {
    console.log(`Toggling enabled from ${enabled} => ${!enabled}`);
    setEnabled(!enabled);
  }
  return (
    <NewNoteContext.Provider value={{ enabled: enabled, toggle: toggle }} >
      {children}
    </NewNoteContext.Provider>
  )
}