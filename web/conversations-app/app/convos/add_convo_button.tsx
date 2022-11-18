'use client';
import { Dialog } from '@headlessui/react';
import { useState } from 'react';

import FabButton from "../../components/fab_button";

export default function AddConvoButton() {
  const [open, setIsOpen] = useState(false);

  return (
    <>
      {/* TODO(GikuyuNderitu) Abstract Dialog into its own component for this project */}
      <Dialog open={open} onClose={() => setIsOpen(false)}>
        <div className="fixed inset-0 bg-black/30" aria-hidden="true" />
        <div className="fixed inset-0 flex items-center justify-center p-4">
          <Dialog.Panel className="bg-on-surface max-w-prose mx-auto rounded-md">
            <Dialog.Title className="p-4 text-on-surface-text text-xl">Add a Conversation</Dialog.Title>
            <Dialog.Description className="px-4 text-on-surface-text">
              This will permanently deactivate your account
            </Dialog.Description>

            <p className="px-4 text-on-surface-text">
              Are you sure you want to deactivate your account? All of your data
              will be permanently removed. This action cannot be undone.
            </p>
            <div className="p-4 flex flex-row-reverse gap-y-3">
              {/* TODO(GikuyuNderitu) Abstract button into its own component for this project */}
              <button className="ml-4 px-4 bg-on-surface-btn text-on-surface-btn-text min-h-tap-target min-w-tap-target rounded-md" onClick={() => setIsOpen(false)}>Submit</button>
              <button className="bg-on-surface-secondary-btn text-on-surface-secondary-btn-text rounded-md" onClick={() => setIsOpen(false)}>Cancel</button>

            </div>
          </Dialog.Panel>
        </div>
      </Dialog>
      <FabButton onClick={() => setIsOpen(true)} />
    </>
  );
}