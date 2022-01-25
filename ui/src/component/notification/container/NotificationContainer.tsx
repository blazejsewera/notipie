import { FC } from '../../../type/react'
import { Notification } from '../../../type/notification'
import { Intl } from '../../../i18l/intl'
import { cx } from '../../../util/cx'
import { Header } from './section/Header'
import { NotificationCardList } from './section/NotificationCardList'
import { NotificationContainerHandlers } from '../../../type/handler'

export interface NotificationContainerProps {
  title: string
  notifications: Notification[]
  handlers: NotificationContainerHandlers
  intl: Intl
}

export const NotificationContainer: FC<NotificationContainerProps> = ({ title, notifications, handlers, intl }) => {
  const { onCheckAll } = handlers
  return (
    <div className={cx('bg-gray-100', 'dark:bg-gray-700', 'inline-block', 'p-5', 'pb-6', 'rounded-3xl', 'shadow-xl')}>
      <Header title={title} onCheckAll={onCheckAll} />
      <NotificationCardList {...{ notifications, intl }} />
    </div>
  )
}
