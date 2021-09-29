import * as React from 'react'
import { cx } from '../../../../src/utils/cx'
import { Footer } from './section/Footer'
import { Header } from './section/Header'
import { Indicator } from './sprite/icon/Indicator'
import { Body } from './text/Body'

export interface NotificationCardProps {
  appName: string
  timestamp: string
  appImgUri?: string
  title: string
  subtitle?: string
  body?: string
}

export const NotificationCard: React.FC<NotificationCardProps> = ({
  appName,
  timestamp,
  appImgUri,
  title,
  subtitle,
  body,
}) => (
  <div className={cx('rounded-3xl', 'bg-white', 'w-80', 'p-5', 'shadow-md', 'relative')}>
    <Header {...{ appName, appImgUri, title, subtitle }} bgColor="darkgreen" />
    <Indicator />
    <Body>{body}</Body>
    <Footer {...{ appName, timestamp }} />
  </div>
)
