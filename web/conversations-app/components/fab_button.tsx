'use client'

import { PlusIcon } from '@heroicons/react/20/solid'

type FabButtonProps = {
  onClick(event: React.MouseEvent<HTMLButtonElement>): void;
}

export default function FabButton({ onClick }: FabButtonProps) {
  return <button className="
    group
    fixed bottom-0 right-0 mb-16 mr-16 p-2
    bg-stone-50 
    rounded-[50%] hover:rounded-lg
    transition-border-radius duration-250 ease-in"
    onClick={onClick}>
    <PlusIcon className="
      w-8 h-8 
      text-blue-500 stroke-blue-500 stroke-0 group-hover:stroke-1
      transition-color duration-250" />
  </button>
}