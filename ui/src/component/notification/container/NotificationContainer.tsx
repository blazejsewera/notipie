import { FC } from '../../../type/react'
import { NotificationWithHandlers } from '../../../type/notification'
import { Intl } from '../../../i18l/intl'
import { cx } from '../../../util/classname/cx'
import { Header } from './section/Header'
import { NotificationCardList } from './section/NotificationCardList'
import { NotificationContainerHandlers } from '../../../type/handler'

export interface NotificationContainerProps {
  title: string
  notificationsWithHandlers: NotificationWithHandlers[]
  handlers: NotificationContainerHandlers
  intl: Intl
  style?: string
}

export const NotificationContainer: FC<NotificationContainerProps> = ({
  title,
  notificationsWithHandlers,
  handlers,
  intl,
  style,
}) => {
  const { onCheckAll } = handlers
  return (
    <div
      className={cx(
        'bg-gray-100',
        'dark:bg-gray-700',
        'inline-block',
        'p-5',
        'pb-6',
        'rounded-3xl',
        'shadow-xl',
        style ?? '',
      )}
    >
      <Header title={title} onCheckAll={onCheckAll} />
      <NotificationCardList {...{ notificationsWithHandlers, intl }} />
    </div>
  )
}
