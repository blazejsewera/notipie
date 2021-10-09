import * as React from 'react'
import type { FC } from '../../../type/react'
import type { Notification } from '../../../type/notification'
import { NotificationCard } from '../card/NotificationCard'
import { cx } from '../../../util/cx'

export interface NotificationContainerProps {
  notifications: Notification[]
}

export const NotificationContainer: FC<NotificationContainerProps> = ({ notifications }) => {
  return (
    <div className={cx('grid grid-cols-1 space-y-5')}>
      {notifications.map((notification) => (
        <NotificationCard key={notification.uuid} notification={notification} />
      ))}
    </div>
  )
}
