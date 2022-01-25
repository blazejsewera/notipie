import * as React from 'react'
import { FC } from '../../../../../type/react'
import { Icon } from '../../../../../external/icon'
import { cx } from '../../../../../util/cx'
import {
  OnNotificationCardArchive,
  OnNotificationCardCheck,
  OnNotificationCardSettings,
} from '../../../../../type/handler'

export interface ControlButtonProps {
  SvgIcon: Icon
  onClick: OnNotificationCardArchive | OnNotificationCardCheck | OnNotificationCardSettings
}

export const ControlButton: FC<ControlButtonProps> = ({ SvgIcon, onClick }) => {
  return (
    <button className="inline" type="button" onClick={onClick}>
      <SvgIcon className={cx('stroke-current', 'text-gray-500', 'dark:text-gray-400', 'w-5', 'h-5')} />
    </button>
  )
}
