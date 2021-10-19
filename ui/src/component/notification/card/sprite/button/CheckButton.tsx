import * as React from 'react'
import { Archive } from '../../../../../external/icon'
import type { OnNotificationCardCheck } from '../../../../../type/handler'
import type { FC } from '../../../../../type/react'
import { ControlButton } from './ControlButton'

export interface CheckButtonProps {
  onClick: OnNotificationCardCheck
}

export const ArchiveButton: FC<CheckButtonProps> = ({ onClick }) => (
  <ControlButton SvgIcon={Archive} onClick={onClick} />
)
