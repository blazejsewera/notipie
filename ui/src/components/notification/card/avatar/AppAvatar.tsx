import React from 'react'
import cx from '../../../../utils/cx'
import ImgAppAvatar from './ImgAppAvatar'
import InitialsAppAvatar from './InitialsAppAvatar'

export interface AppAvatarProps {
  appName: string
  imgUri?: string
  bgColor?: string
  size?: 'small' | 'medium' | 'large'
}

const AppAvatar: React.FC<AppAvatarProps> = ({ appName, imgUri, bgColor, size = 'medium' }) => {
  const sizeClassesDict = {
    small: ['h-6', 'w-6'],
    medium: ['h-8', 'w-8'],
    large: ['h-10', 'w-10'],
  }
  const sizeClasses = sizeClassesDict[size]
  const shapeClasses = ['rounded-full', 'overflow-clip', 'overflow-hidden']

  if (imgUri) {
    return (
      <div className={cx(...sizeClasses, ...shapeClasses)}>
        <ImgAppAvatar imgUri={imgUri} />
      </div>
    )
  }

  return (
    <div className={cx(...sizeClasses, ...shapeClasses)}>
      <InitialsAppAvatar {...{ appName, bgColor }} />
    </div>
  )
}

export default AppAvatar
