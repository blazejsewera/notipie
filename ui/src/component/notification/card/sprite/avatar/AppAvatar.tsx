import * as React from 'react'
import type { FC } from '../../../../../type/react'
import { cx } from '../../../../../util/cx'
import { ImgAppAvatar } from './ImgAppAvatar'
import { InitialsAppAvatar } from './InitialsAppAvatar'

export interface AppAvatarProps {
  appName: string
  appImgUri?: string
  bgColor?: string
  size?: 'small' | 'medium' | 'large'
}

export const AppAvatar: FC<AppAvatarProps> = ({ appName, appImgUri, bgColor, size = 'medium' }) => {
  const sizeClasses = {
    small: ['h-8', 'w-8'],
    medium: ['h-12', 'w-12'],
    large: ['h-14', 'w-14'],
  }[size]
  const styleClasses = ['rounded-full', 'overflow-clip', 'overflow-hidden', 'inline-block', 'cursor-default']

  if (appImgUri) {
    return (
      <div className={cx(...sizeClasses, ...styleClasses)}>
        <ImgAppAvatar appImgUri={appImgUri} />
      </div>
    )
  }

  return (
    <div className={cx(...sizeClasses, ...styleClasses)}>
      <InitialsAppAvatar {...{ appName, bgColor }} />
    </div>
  )
}
