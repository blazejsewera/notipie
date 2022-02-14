import { FC } from '../../../../type/react'
import { Archive, Check, ExternalLink } from '../../../../external/icon'
import { ControlButton } from '../sprite/button/ControlButton'
import { NotificationCardHandlers } from '../../../../type/handler'

export interface ControlsProps {
  handlers: NotificationCardHandlers
}

export const Controls: FC<ControlsProps> = ({ handlers }) => (
  <div className="inline-grid grid-cols-3 gap-x-5 ml-auto my-auto mr-2">
    <ControlButton SvgIcon={Archive} onClick={handlers.onArchive} />
    <ControlButton SvgIcon={Check} onClick={handlers.onCheck} />
    <ControlButton SvgIcon={ExternalLink} onClick={handlers.onExternal} />
  </div>
)
