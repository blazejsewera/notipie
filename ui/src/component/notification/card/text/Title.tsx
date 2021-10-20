import * as React from 'react'
import type { FC } from '../../../../type/react'
import { cx } from '../../../../util/cx'

export interface TitleProps {
  children: string
}

export const Title: FC<TitleProps> = ({ children }) => (
  <p className={cx('text-lg', 'font-bold', 'mb-0', 'leading-5', 'text-gray-800', 'dark:text-gray-300')}>{children}</p>
)
