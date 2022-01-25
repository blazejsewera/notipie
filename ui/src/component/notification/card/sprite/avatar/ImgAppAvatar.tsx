import { FC } from '../../../../../type/react'

export interface ImgAppAvatarProps {
  appImgUri: string
}

export const ImgAppAvatar: FC<ImgAppAvatarProps> = ({ appImgUri }) => <img src={appImgUri} />
