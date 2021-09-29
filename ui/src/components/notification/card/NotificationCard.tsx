import * as React from 'react'
import { AppAvatar } from './avatar/AppAvatar'
import { Subtitle } from './text/Subtitle'
import { Title } from './text/Title'
import { Body } from './text/Body'
import { Footnote } from './text/Footnote'

export interface NotificationCardProps {
  appName: string
  timestamp: string
  appImgUri?: string
  title: string
  subtitle?: string
  body?: string
}

export const NotificationCard: React.FC<NotificationCardProps> = ({
  appName,
  timestamp,
  appImgUri,
  title,
  subtitle,
  body,
}) => (
  <div>
    <AppAvatar appName={appName} imgUri={appImgUri} />
    <Title>{title}</Title>
    <Subtitle>{subtitle}</Subtitle>
    <Body>{body}</Body>
    <Footnote {...{ appName, timestamp }} />
  </div>
)
