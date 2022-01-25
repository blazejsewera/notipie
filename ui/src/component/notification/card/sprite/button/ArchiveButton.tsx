import { Archive } from '../../../../../external/icon'
import { OnNotificationCardArchive } from '../../../../../type/handler'
import { FC } from '../../../../../type/react'
import { ControlButton } from './ControlButton'

export interface ArchiveButtonProps {
  onClick: OnNotificationCardArchive
}

export const ArchiveButton: FC<ArchiveButtonProps> = ({ onClick }) => (
  <ControlButton SvgIcon={Archive} onClick={onClick} />
)
