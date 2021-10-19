import * as React from 'react'
import type { FC } from '../../../../type/react'
import type { Notification } from '../../../../type/notification'
import type { Intl } from '../../../../i18l/intl'
import { cx } from '../../../../util/cx'
import { NotificationCard } from '../../card/NotificationCard'

export interface NotificationCardListProps {
  notifications: Notification[]
  intl: Intl
}

export const NotificationCardList: FC<NotificationCardListProps> = ({ notifications, intl }) => (
  <div className={cx('grid grid-cols-1 space-y-5')}>
    {notifications.map((notification) => (
      <NotificationCard key={notification.id} {...{ notification, intl }} />
    ))}
  </div>
)
