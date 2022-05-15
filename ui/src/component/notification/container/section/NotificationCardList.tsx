import { FC } from '../../../../type/react'
import { Notification } from '../../../../type/notification'
import { Intl } from '../../../../i18l/intl'
import { cx } from '../../../../util/classname/cx'
import { NotificationCard } from '../../card/NotificationCard'

export interface NotificationCardListProps {
  notifications: Notification[]
  intl: Intl
}

export const NotificationCardList: FC<NotificationCardListProps> = ({ notifications, intl }) => (
  <div className={cx('grid grid-cols-1 space-y-5')}>
    {notifications.map(notification => (
      <NotificationCard key={notification.id} {...{ notification, intl }} />
    ))}
  </div>
)
