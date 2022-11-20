'use client';
import { Dialog } from '@headlessui/react';
import { useState } from 'react';

import FabButton from '../../components/fab_button';
import FlatButton from '../../components/flat_button';

// TODO(GikuyuNderitu) Wrap this some sort of Tutorial service to inform the
// what this button does. 
export default function AddConvoButton() {
  const [open, setIsOpen] = useState(false);

  return (
    <>
      {
      /* TODO(GikuyuNderitu) Abstract Dialog into its own component for this
       * project */}
      <Dialog open={open} onClose={() => setIsOpen(false)}>
        <div className="fixed inset-0 bg-black/30" aria-hidden="true" />
        <div className="fixed inset-0 flex items-center justify-center p-4">
          <Dialog.Panel className="bg-on-surface max-w-prose mx-auto rounded-md">
            <Dialog.Title className="p-4 text-on-surface-text text-xl">
              Add a Conversation
            </Dialog.Title>
            <Dialog.Description className="px-4 text-on-surface-text">
              Add a title to get started!
            </Dialog.Description>

            <p className="px-4 text-on-surface-text">
              Are you sure you want to deactivate your account? All of your
              data will be permanently removed. This action cannot be undone.
            </p>
            <div className="p-4 flex flex-row-reverse gap-y-3">
              <FlatButton
                className="ml-4 bg-on-surface-btn text-on-surface-btn-text"
                onClick={() => setIsOpen(false)}>Submit</FlatButton>
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