'use client';

import { useContext } from "react";
import { NewNoteContext } from "./notes_context";
import FabButton from "../../../components/fab_button";

export default function AddNoteButton() {
  const newNoteContext = useContext(NewNoteContext);
  return <FabButton onClick={newNoteContext.toggle} />
}