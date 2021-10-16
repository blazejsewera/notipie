import * as React from 'react'
import type { FC } from '../../../../../type/react'
import { cx } from '../../../../../util/cx'

export const Indicator: FC = () => (
  <div className={cx('rounded-full', 'w-1.5', 'h-1.5', 'bg-yellow-500', 'absolute', 'top-4', 'right-4')}></div>
)