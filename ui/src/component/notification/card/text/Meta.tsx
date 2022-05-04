import { FC } from '../../../../type/react'
import { cx } from '../../../../util/classname/cx'

export interface MetaProps {
  appName: string
  timestamp: string
  relativeTime?: string
}

export const Meta: FC<MetaProps> = ({ appName, relativeTime, timestamp }) => {
  const textClasses = ['text-xs', 'font-bold', 'text-gray-500', 'dark:text-gray-400']
  return (
    <div className="inline-block">
      <p className={cx(...textClasses)}>— by {appName}</p>
      <p className={cx(...textClasses)}>{relativeTime ? relativeTime : timestamp}</p>
    </div>
  )
}
