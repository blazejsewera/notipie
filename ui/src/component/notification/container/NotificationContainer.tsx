import * as React from 'react'
import type { FC } from '../../../type/react'
import type { Notification } from '../../../type/notification'
import type { Intl } from '../../../i18l/intl'
import { cx } from '../../../util/cx'
import { Header } from './section/Header'
import { NotificationCardList } from './section/NotificationCardList'

export interface NotificationContainerProps {
  title: string
  notifications: Notification[]
  intl: Intl
}

export const NotificationContainer: FC<NotificationContainerProps> = ({ title, notifications, intl }) => {
  return (
    <div className={cx('bg-gray-100', 'inline-block', 'p-5', 'pb-6', 'rounded-3xl', 'shadow-xl')}>
      <Header title={title} />
      <NotificationCardList {...{ notifications, intl }} />
    </div>
  )
}
