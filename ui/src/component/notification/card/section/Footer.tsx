import { NotificationCardHandlers } from '../../../../type/handler'
import { FC } from '../../../../type/react'
import { cx } from '../../../../util/cx'
import { Meta } from '../text/Meta'
import { Controls } from './Controls'

export interface FooterProps {
  appName: string
  timestamp: string
  handlers: NotificationCardHandlers
}

export const Footer: FC<FooterProps> = ({ appName, timestamp, handlers }) => (
  <div className={cx('flex')}>
    <Meta {...{ appName, timestamp }} />
    <Controls handlers={handlers} />
  </div>
)
