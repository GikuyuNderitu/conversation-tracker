'use client';

import { useAtom } from "jotai";
import FabButton from "../../../components/fab_button";
import { newNoteAtom } from "./new_note_state";

export default function AddNoteButton() {
  const [enabled, setEnabled] = useAtom(newNoteAtom);
  return <FabButton onClick={() => setEnabled(!enabled)} />
}