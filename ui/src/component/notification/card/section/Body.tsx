import { Intl } from '../../../../i18l/intl'
import { FC } from '../../../../type/react'
import { cx } from '../../../../util/cx'
import { TrimmableText as Trim } from '../text/TrimmableText'

export interface BodyProps {
  trimThreshold?: number
  intl: Intl
  children: string
}

export const Body: FC<BodyProps> = ({ intl, children, trimThreshold }) => {
  const text = children

  return (
    <p className={cx('whitespace-pre-line', 'text-sm', 'leading-4', 'my-3', 'text-gray-800', 'dark:text-gray-300')}>
      <Trim threshold={trimThreshold} intl={intl}>
        {text}
      </Trim>
    </p>
  )
}
