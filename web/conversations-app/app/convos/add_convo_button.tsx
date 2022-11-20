'use client';
import { Dialog } from '@headlessui/react';
import { useRouter } from "next/navigation";
import { useState } from 'react';

import FabButton from '../../components/fab_button';
import FlatButton from '../../components/flat_button';

async function create(title: string, refresh: () => void) {
  // TODO(GikuyuNderitu): Handle this if there's an error status
  await fetch(`/api/convos`, {
    method: 'POST',
    headers: {
      'Content-type': 'application/json',
    },
    body: JSON.stringify({ title }),
  });

  refresh();
}

// TODO(GikuyuNderitu) Wrap this some sort of Tutorial service to inform the
// what this button does. 
export default function AddConvoButton() {
  const [open, setIsOpen] = useState(false);
  const [title, setTitle] = useState("");
  const router = useRouter();

  function onChange(e: React.ChangeEvent<HTMLInputElement>) {
    setTitle(e.target.value);
    console.log(e.target.value);
  }

  async function addConversation() {
    // TODO(GikuyuNderitu): Add basic validation for empty title. Display an 
    // Error message if title is empty.

    // TODO(GikuyuNderitu): Create an api endpoint and call to create conversation.
    await create(title, router.refresh);

    setIsOpen(false);
  }

  return (
    <>
      {
        /** 
         * TODO(GikuyuNderitu) Abstract Dialog into its own component for this
         * project 
         */
      }
      <Dialog open={open} onClose={() => setIsOpen(false)}>
        <div className="fixed inset-0 bg-black/30" aria-hidden="true" />
        <div className="fixed inset-0 flex items-center justify-center p-4">
          <Dialog.Panel className="bg-on-surface max-w-prose mx-auto rounded-md min-w-[280px]">
            <Dialog.Title className="p-4 text-on-surface-text text-xl">
              Add a Conversation
            </Dialog.Title>

            <div className="px-4">
              <input className="text-black rounded w-full" type="text" value={title} onChange={onChange} />
            </div>

            <div className="p-4 flex flex-row-reverse gap-y-3">
              <FlatButton
                className="ml-4 bg-on-surface-btn text-on-surface-btn-text"
                onClick={addConversation}>Add Convo</FlatButton>
              <FlatButton
                className="
                  bg-on-surface-secondary-btn
                  text-on-surface-secondary-btn-text"
                onClick={() => setIsOpen(false)}>Cancel</FlatButton>
            </div>
          </Dialog.Panel>
        </div>
      </Dialog>
      <FabButton onClick={() => setIsOpen(true)} />
    </>
  );
}