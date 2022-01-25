import { FC } from '../../../type/react'
import { Notification } from '../../../type/notification'
import { Intl } from '../../../i18l/intl'
import { cx } from '../../../util/cx'
import { Header } from './section/Header'
import { Body } from './section/Body'
import { Footer } from './section/Footer'
import { Indicator } from './sprite/icon/Indicator'

export interface NotificationCardProps {
  notification: Notification
  intl: Intl
}

export const NotificationCard: FC<NotificationCardProps> = ({ notification, intl }) => {
  const { appName, appImgUri, title, subtitle, body, timestamp, handlers } = notification

  return (
    <div className={cx('rounded-3xl', 'bg-white', 'dark:bg-gray-800', 'w-80', 'p-5', 'shadow-lg', 'relative')}>
      <Header {...{ appName, appImgUri, title, subtitle }} />
      <Indicator />
      <Body intl={intl}>{body ?? ''}</Body>
      <Footer {...{ appName, timestamp, handlers }} />
    </div>
  )
}
