import { FC } from '../../../../type/react'
import { NotificationWithHandlers } from '../../../../type/notification'
import { Intl } from '../../../../i18l/intl'
import { cx } from '../../../../util/classname/cx'
import { NotificationCard } from '../../card/NotificationCard'

export interface NotificationCardListProps {
  notificationsWithHandlers: NotificationWithHandlers[]
  intl: Intl
}

export const NotificationCardList: FC<NotificationCardListProps> = ({ notificationsWithHandlers, intl }) => (
  <div className={cx('grid grid-cols-1 space-y-5')}>
    {notificationsWithHandlers.map(notificationWithHandlers => (
      <NotificationCard key={notificationWithHandlers.notification.id} {...{ notificationWithHandlers, intl }} />
    ))}
  </div>
)
