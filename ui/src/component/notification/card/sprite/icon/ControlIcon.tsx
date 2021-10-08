import * as React from 'react'
import type { FC } from '../../../../../type/react'
import type { Icon } from 'react-feather'
import { cx } from '../../../../../util/cx'

export interface ControlIconProps {
  SvgIcon: Icon
}

export const ControlIcon: FC<ControlIconProps> = ({ SvgIcon }) => {
  return (
    <div className="inline">
      <SvgIcon className={cx('stroke-current', 'text-gray-500', 'w-5', 'h-5')} />
    </div>
  )
}
