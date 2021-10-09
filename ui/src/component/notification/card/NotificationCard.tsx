import * as React from 'react'
import type { FC } from '../../../type/react'
import type { Notification } from '../../../type/notification'
import { cx } from '../../../util/cx'
import { Footer } from './section/Footer'
import { Header } from './section/Header'
import { Indicator } from './sprite/icon/Indicator'
import { Body } from './text/Body'

export interface NotificationCardProps {
  notification: Notification
}

export const NotificationCard: FC<NotificationCardProps> = ({ notification }) => {
  const { appName, appImgUri, title, subtitle, body, timestamp } = notification

  return (
    <div className={cx('rounded-3xl', 'bg-white', 'w-80', 'p-5', 'shadow-md', 'relative')}>
      <Header {...{ appName, appImgUri, title, subtitle }} bgColor="darkgreen" />
      <Indicator />
      {body ? <Body>{body}</Body> : null}
      <Footer {...{ appName, timestamp }} />
    </div>
  )
}
