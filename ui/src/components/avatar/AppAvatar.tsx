import React from 'react'
import cx from 'src/utils/cx'
import ImgAppAvatar from './ImgAppAvatar'
import InitialsAppAvatar from './InitialsAppAvatar'

export interface AppAvatarProps {
  appName: string
  img?: string
  size?: 'small' | 'medium' | 'large'
}

const AppAvatar: React.FC<AppAvatarProps> = ({ appName, img, size = 'medium' }) => {
  let sizeClasses: string[]
  switch (size) {
    case 'small':
      sizeClasses = ['', '']
      break
    case 'large':
      sizeClasses = ['', '']
      break
    default:
      sizeClasses = ['', '']
      break
  }

  if (img) {
    return (
      <div className={cx(...sizeClasses)}>
        <ImgAppAvatar imgPath={img} />
      </div>
    )
  }

  return (
    <div className={cx(...sizeClasses)}>
      <InitialsAppAvatar initials={makeInitials(appName)} />
    </div>
  )
}

const makeInitials = (name: string) => name.substr(0, 2)

export default AppAvatar
