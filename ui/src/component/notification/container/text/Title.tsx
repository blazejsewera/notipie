import { FC } from '../../../../type/react'
import { cx } from '../../../../util/classname/cx'

export interface TitleProps {
  children: string
}

export const Title: FC<TitleProps> = ({ children }) => (
  <p className={cx('text-2xl', 'font-bold', 'text-gray-800', 'dark:text-gray-300')}>{children}</p>
)
