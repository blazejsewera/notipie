import * as React from 'react'
import type { FC } from '../../../../type/react'
import { cx } from '../../../../util/cx'

export interface MetaProps {
  appName: string
  timestamp: string
}

export const Meta: FC<MetaProps> = ({ appName, timestamp }) => {
  const textClasses = ['text-xs', 'font-bold', 'text-gray-500', 'dark:text-gray-400']
  return (
    <div className="inline-block">
      <p className={cx(...textClasses)}>— by {appName}</p>
      <p className={cx(...textClasses)}>{timestamp}</p>
    </div>
  )
}
