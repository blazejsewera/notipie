import * as React from 'react'
import { Archive } from '../../../../../external/icon'
import type { OnNotificationCardSettings } from '../../../../../type/handler'
import type { FC } from '../../../../../type/react'
import { ControlButton } from './ControlButton'

export interface SettingsButtonProps {
  onClick: OnNotificationCardSettings
}

export const SettingsButton: FC<SettingsButtonProps> = ({ onClick }) => (
  <ControlButton SvgIcon={Archive} onClick={onClick} />
)
