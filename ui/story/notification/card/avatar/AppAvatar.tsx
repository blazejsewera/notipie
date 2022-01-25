import * as React from 'react'
import { Meta, Story } from '@storybook/react'
import { AppAvatarProps } from '../../../../src/component/notification/card/sprite/avatar/AppAvatar'
import { AppAvatar } from '../../../../src/component/notification/card/sprite/avatar/AppAvatar'

export default {
  title: 'notification/card/avatar/AppAvatar',
  component: AppAvatar,
} as Meta

export const ExampleAppAvatar: Story<AppAvatarProps> = args => <AppAvatar {...args} />

ExampleAppAvatar.args = {
  appName: 'Testapp',
  size: 'medium',
}
