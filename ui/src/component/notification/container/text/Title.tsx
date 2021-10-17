import * as React from 'react'
import type { FC } from '../../../../type/react'
import { cx } from '../../../../util/cx'

export interface TitleProps {
  children: string
}

export const Title: FC<TitleProps> = ({ children }) => (
  <p className={cx('text-2xl', 'font-bold', 'text-gray-800')}>{children}</p>
)
