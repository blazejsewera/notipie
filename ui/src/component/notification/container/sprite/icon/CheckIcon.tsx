import * as React from 'react'
import type { FC } from '../../../../../type/react'
import { CheckSquare } from '../../../../../external/icon'
import { cx } from '../../../../../util/cx'

export const CheckIcon: FC = () => (
  <CheckSquare className={cx('stroke-current', 'text-gray-500', 'w-5', 'h-5', 'ml-auto', 'my-auto')} />
)
