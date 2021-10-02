import * as React from 'react'
import { useState } from 'react'
import { cx } from '../../../../util/cx'

export interface BodyProps {
  trimThreshold: number
}

// TODO: Add tests to this component
export const Body: React.FC<BodyProps> = ({ children, trimThreshold = 150 }) => {
  const text = children as string

  const styleClasses = ['whitespace-pre-line', 'text-sm', 'leading-4', 'my-3', 'text-gray-800']

  if (text.length <= trimThreshold) {
    return <p className={cx(...styleClasses)}>{text}</p>
  }

  const [isTrimmed, setIsTrimmed] = useState(true)
  return (
    <p className={cx(...styleClasses)}>
      {isTrimmed ? text.slice(0, trimThreshold) : text}
      <span className="text-blue-500 cursor-pointer" onClick={() => setIsTrimmed(!isTrimmed)}>
        {isTrimmed ? '[...]' : ' [show less]'}
      </span>
    </p>
  )
}
