import * as React from 'react'

export interface ImgAppAvatarProps {
  appImgUri: string
}

export const ImgAppAvatar: React.FC<ImgAppAvatarProps> = ({ appImgUri }) => <img src={appImgUri} />
