import * as React from 'react'
import type { FC } from '../../../../type/react'
import { Archive, Check, Settings } from '../../../../external/icon'
import { ControlButton } from '../sprite/button/ControlButton'
import type { NotificationCardHandlers } from '../../../../type/handler'

export interface ControlsProps {
  handlers: NotificationCardHandlers
}

export const Controls: FC<ControlsProps> = ({ handlers }) => (
  <div className="inline-grid grid-cols-3 gap-x-5 ml-auto my-auto mr-2">
    <ControlButton SvgIcon={Archive} onClick={handlers.onArchive} />
    <ControlButton SvgIcon={Check} onClick={handlers.onCheck} />
    <ControlButton SvgIcon={Settings} onClick={handlers.onSettings} />
  </div>
)
