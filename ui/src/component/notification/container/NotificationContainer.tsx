import * as React from 'react'
import type { FC } from '../../../type/react'
import type { Notification } from '../../../type/notification'
import type { Intl } from '../../../i18l/intl'
import { NotificationCard } from '../card/NotificationCard'
import { cx } from '../../../util/cx'
import { Header } from './section/Header'

export interface NotificationContainerProps {
  title: string
  notifications: Notification[]
  intl: Intl
}

export const NotificationContainer: FC<NotificationContainerProps> = ({ title, notifications, intl }) => {
  return (
    <div className={cx('bg-gray-100', 'inline-block', 'p-5', 'rounded-xl', 'shadow-lg')}>
      <Header title={title} />
      <div className={cx('grid grid-cols-1 space-y-5')}>
        {notifications.map((notification) => (
          <NotificationCard key={notification.id} notification={notification} intl={intl} />
        ))}
      </div>
    </div>
  )
}
