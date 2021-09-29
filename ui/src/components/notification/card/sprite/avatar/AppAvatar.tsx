import * as React from 'react'
import { cx } from '../../../../../utils/cx'
import { ImgAppAvatar } from './ImgAppAvatar'
import { InitialsAppAvatar } from './InitialsAppAvatar'

export interface AppAvatarProps {
  appName: string
  appImgUri?: string
  bgColor?: string
  size?: 'small' | 'medium' | 'large'
}

export const AppAvatar: React.FC<AppAvatarProps> = ({ appName, appImgUri, bgColor, size = 'medium' }) => {
  const sizeClasses = {
    small: ['h-8', 'w-8'],
    medium: ['h-10', 'w-10'],
    large: ['h-12', 'w-12'],
  }[size]
  const shapeClasses = ['rounded-full', 'overflow-clip', 'overflow-hidden']
  const displayClass = 'inline-block'

  if (appImgUri) {
    return (
      <div className={cx(...sizeClasses, ...shapeClasses, displayClass)}>
        <ImgAppAvatar appImgUri={appImgUri} />
      </div>
    )
  }

  return (
    <div className={cx(...sizeClasses, ...shapeClasses, displayClass)}>
      <InitialsAppAvatar {...{ appName, bgColor }} />
    </div>
  )
}
