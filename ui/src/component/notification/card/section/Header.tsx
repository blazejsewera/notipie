import * as React from 'react'
import { cx } from '../../../../util/cx'
import { AppAvatar } from '../sprite/avatar/AppAvatar'
import { Subtitle } from '../text/Subtitle'
import { Title } from '../text/Title'

export interface HeaderProps {
  appName: string
  appImgUri?: string
  bgColor?: string
  title: string
  subtitle?: string
}

export const Header: React.FC<HeaderProps> = ({ appName, appImgUri, bgColor, title, subtitle }) => (
  <div className={cx('flex')}>
    <AppAvatar size="medium" {...{ appName, appImgUri, bgColor }} />
    <div className={cx('my-auto', 'inline-block', 'ml-5')}>
      <Title>{title}</Title>
      <Subtitle>{subtitle}</Subtitle>
    </div>
  </div>
)
