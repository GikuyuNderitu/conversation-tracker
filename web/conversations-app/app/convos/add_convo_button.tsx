'use client';
import FabButton from "../../components/fab_button";

export default function AddConvoButton() {

  return (
    <FabButton onClick={() => console.log('adding conversation')} />
  );
}