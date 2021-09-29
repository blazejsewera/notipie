import * as React from 'react'
import { cx } from '../../../../utils/cx'

export interface MetaProps {
  appName: string
  timestamp: string
}

export const Meta: React.FC<MetaProps> = ({ appName, timestamp }) => {
  const textClasses = ['text-xs', 'font-bold', 'text-gray-500']
  return (
    <div className="inline-block">
      <p className={cx(...textClasses)}>— by {appName}</p>
      <p className={cx(...textClasses)}>{timestamp}</p>
    </div>
  )
}
