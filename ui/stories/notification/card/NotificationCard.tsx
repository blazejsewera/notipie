import * as React from 'react'
import { Story, Meta } from '@storybook/react'

import { NotificationCard, NotificationCardProps } from '../../../src/components/notification/card/NotificationCard'
import { fiveSentenceLoremIpsum } from '../../assets/lipsum'

export default {
  title: 'notification/card/NotificationCard',
  component: NotificationCard,
} as Meta

export const ExampleNotificationCard: Story<NotificationCardProps> = (args) => <NotificationCard {...args} />

ExampleNotificationCard.args = {
  appName: 'Example App',
  timestamp: '2 hours ago',
  title: 'Example Title',
  subtitle: 'Example Subtitle',
  body: fiveSentenceLoremIpsum,
}
