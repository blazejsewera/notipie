import { Archive } from '../../../../../external/icon'
import { OnNotificationCardCheck } from '../../../../../type/handler'
import { FC } from '../../../../../type/react'
import { ControlButton } from './ControlButton'

export interface CheckButtonProps {
  onClick: OnNotificationCardCheck
}

export const ArchiveButton: FC<CheckButtonProps> = ({ onClick }) => (
  <ControlButton SvgIcon={Archive} onClick={onClick} />
)
