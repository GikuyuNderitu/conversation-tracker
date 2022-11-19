import React from "react"

type OnClickHandler<T extends HTMLElement> = (event: React.MouseEvent<T>) => void;
const _flatButtonStyle = "min-h-tap-target min-w-tap-target rounded-md px-4";

type FlatButtonProps = {
  onClick: OnClickHandler<HTMLButtonElement> | undefined;
  className: string | undefined;
  children?: React.ReactNode;
}

export default function FlatButton({ onClick, className, children }: FlatButtonProps) {
  return (
    <button className={getClassName(className)} onClick={onClick}>{children}</button>
  );
}

function getClassName(className: string | undefined): string {
  if (className == undefined) return _flatButtonStyle
  return `${_flatButtonStyle} ${className}`;
}