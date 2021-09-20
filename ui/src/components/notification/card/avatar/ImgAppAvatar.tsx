import React from 'react'

export interface ImgAppAvatarProps {
  imgPath: string
}

const ImgAppAvatar: React.FC<ImgAppAvatarProps> = ({ imgPath }) => <img src={imgPath} />

export default ImgAppAvatar
