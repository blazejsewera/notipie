import { Archive } from '../../../../../external/icon'
import { OnNotificationCardSettings } from '../../../../../type/handler'
import { FC } from '../../../../../type/react'
import { ControlButton } from './ControlButton'

export interface SettingsButtonProps {
  onClick: OnNotificationCardSettings
}

export const SettingsButton: FC<SettingsButtonProps> = ({ onClick }) => (
  <ControlButton SvgIcon={Archive} onClick={onClick} />
)
