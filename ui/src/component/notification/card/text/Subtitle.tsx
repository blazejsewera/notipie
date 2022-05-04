import { FC } from '../../../../type/react'
import { cx } from '../../../../util/classname/cx'

export interface SubtitleProps {
  children: string
}

export const Subtitle: FC<SubtitleProps> = ({ children }) => (
  <p className={cx('text-xs', 'font-bold', 'truncate', 'text-gray-500', 'dark:text-gray-400')}>{children}</p>
)
