import * as React from 'react'
import type { FC } from '../../../../type/react'
import { useState } from 'react'
import { cx } from '../../../../util/cx'

export interface BodyProps {
  trimThreshold?: number
  children: string
}

export const Body: FC<BodyProps> = ({ children, trimThreshold = 150 }) => {
  const text = children
  const styleClasses = ['whitespace-pre-line', 'text-sm', 'leading-4', 'my-3', 'text-gray-800']

  if (text.length <= trimThreshold) {
    return <p className={cx(...styleClasses)}>{text}</p>
  }

  const [isTrimmed, setIsTrimmed] = useState(true)
  return (
    <p className={cx(...styleClasses)}>
      {isTrimmed ? text.slice(0, trimThreshold) : text}
      <span
        className="text-blue-500 cursor-pointer"
        onClick={() => setIsTrimmed(!isTrimmed)}
        aria-label={isTrimmed ? 'show more' : 'show less'}
      >
        {isTrimmed ? '[...]' : ' [show less]'}
      </span>
    </p>
  )
}
