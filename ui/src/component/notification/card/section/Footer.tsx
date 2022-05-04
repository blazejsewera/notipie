import { NotificationCardHandlers } from '../../../../type/handler'
import { FC } from '../../../../type/react'
import { cx } from '../../../../util/classname/cx'
import { Meta } from '../text/Meta'
import { Controls } from './Controls'

export interface FooterProps {
  appName: string
  timestamp: string
  relativeTime?: string
  handlers: NotificationCardHandlers
}

export const Footer: FC<FooterProps> = ({ appName, relativeTime, timestamp, handlers }) => (
  <div className={cx('flex')}>
    <Meta {...{ appName, relativeTime, timestamp }} />
    <Controls handlers={handlers} />
  </div>
)
