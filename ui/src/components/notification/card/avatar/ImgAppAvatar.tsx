import * as React from 'react'

export interface ImgAppAvatarProps {
  imgUri: string
}

export const ImgAppAvatar: React.FC<ImgAppAvatarProps> = ({ imgUri }) => <img src={imgUri} />
