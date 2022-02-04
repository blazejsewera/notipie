import { FC } from '../../../../../type/react'
import { cx } from '../../../../../util/classname/cx'

export const Indicator: FC = () => (
  <div
    className={cx(
      'rounded-full',
      'w-2.5',
      'h-2.5',
      'bg-yellow-500',
      'dark:bg-yellow-600',
      'absolute',
      'top-4',
      'right-4',
    )}
  ></div>
)
