import * as React from 'react'
import { useState } from 'react'
import type { Intl } from '../../../../i18l/intl'
import type { FC } from '../../../../type/react'

export interface TrimmableTextProps {
  threshold?: number
  intl: Intl
  children: string
}

export const TrimmableText: FC<TrimmableTextProps> = ({ children, threshold = 150, intl }) => {
  const text = children
  const gt = intl.getTranslation

  if (text.length <= threshold) {
    return <span>{text}</span>
  }

  const [isTrimmed, setIsTrimmed] = useState(true)
  return (
    <span>
      {isTrimmed ? text.slice(0, threshold) : text}
      <span
        className="text-blue-500 dark:text-blue-400 cursor-pointer"
        onClick={() => setIsTrimmed(!isTrimmed)}
        aria-label={isTrimmed ? gt('show more') : gt('show less')}
      >
        {isTrimmed ? gt('[...]') : gt(' [show less]')}
      </span>
    </span>
  )
}
