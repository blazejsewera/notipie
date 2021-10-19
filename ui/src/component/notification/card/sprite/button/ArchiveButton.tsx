import * as React from 'react'
import { Archive } from '../../../../../external/icon'
import type { OnNotificationCardArchive } from '../../../../../type/handler'
import type { FC } from '../../../../../type/react'
import { ControlButton } from './ControlButton'

export interface ArchiveButtonProps {
  onClick: OnNotificationCardArchive
}

export const ArchiveButton: FC<ArchiveButtonProps> = ({ onClick }) => (
  <ControlButton SvgIcon={Archive} onClick={onClick} />
)
