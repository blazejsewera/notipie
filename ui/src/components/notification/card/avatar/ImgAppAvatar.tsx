import React from 'react'

export interface ImgAppAvatarProps {
  imgUri: string
}

const ImgAppAvatar: React.FC<ImgAppAvatarProps> = ({ imgUri }) => <img src={imgUri} />

export default ImgAppAvatar
