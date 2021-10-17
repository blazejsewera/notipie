import * as React from 'react'
import type { FC } from '../../../../type/react'
import { Title } from '../text/Title'
import { CheckSquare } from '../../../../external/icon'
import { cx } from '../../../../util/cx'

export interface HeaderProps {
  title: string
}

export const Header: FC<HeaderProps> = ({ title }) => (
  <div className={cx('flex', 'mb-3', 'mx-2')}>
    <Title>{title}</Title>
    <CheckSquare className={cx('stroke-current', 'text-gray-500', 'w-5', 'h-5', 'ml-auto')} />
  </div>
)
