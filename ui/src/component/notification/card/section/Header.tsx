import { FC } from '../../../../type/react'
import { cx } from '../../../../util/classname/cx'
import { AppAvatar } from '../sprite/avatar/AppAvatar'
import { Subtitle } from '../text/Subtitle'
import { Title } from '../text/Title'

export interface HeaderProps {
  appName: string
  appImgUri?: string
  /**
   * @deprecated bgColor will be generated automatically from appName in the future,
   * expect this prop to be removed
   */
  bgColor?: string
  title: string
  subtitle?: string
}

export const Header: FC<HeaderProps> = ({ appName, appImgUri, bgColor, title, subtitle }) => (
  <div className={cx('flex')}>
    <AppAvatar size="medium" {...{ appName, appImgUri, bgColor }} />
    <div className={cx('my-auto', 'inline-block', 'ml-5')}>
      <Title>{title}</Title>
      {subtitle ? <Subtitle>{subtitle}</Subtitle> : null}
    </div>
  </div>
)
