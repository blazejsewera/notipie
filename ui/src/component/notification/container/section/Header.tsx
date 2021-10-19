import * as React from 'react'
import type { FC } from '../../../../type/react'
import { Title } from '../text/Title'
import { cx } from '../../../../util/cx'
import { CheckIcon } from '../sprite/icon/CheckIcon'

export interface HeaderProps {
  title: string
}

export const Header: FC<HeaderProps> = ({ title }) => (
  <div className={cx('flex', 'mb-4', 'mx-2')}>
    <Title>{title}</Title>
    <CheckIcon />
  </div>
)
