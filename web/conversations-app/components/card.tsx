type CardProps = {
  children?: React.ReactNode;
  className?: string;
}

const defaultClassName = 'min-w-[250px] p-4 bg-on-surface text-on-surface-text rounded-xl'

export default function Card({ className, children }: CardProps) {
  const cardClass = className === undefined ? defaultClassName : defaultClassName.concat(' ', className)
  return (
    <div className={cardClass}>
      {children}
    </div>
  )
}