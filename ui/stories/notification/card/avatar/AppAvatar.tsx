import React from 'react'
import { Story, Meta } from '@storybook/react'

import AppAvatar, { AppAvatarProps } from '../../../../src/components/notification/card/avatar/AppAvatar'

export default {
  title: 'Notification/Card/AppAvatar',
  component: AppAvatar,
} as Meta

export const ExampleAppAvatar: Story<AppAvatarProps> = (args) => <AppAvatar {...args} />

ExampleAppAvatar.args = {
  appName: 'Testapp',
  size: 'medium',
}
