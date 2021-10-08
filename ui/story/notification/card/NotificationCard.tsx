import * as React from 'react'
import { Meta, Story } from '@storybook/react'
import type { NotificationCardProps } from '../../../src/component/notification/card/NotificationCard'
import { NotificationCard } from '../../../src/component/notification/card/NotificationCard'
import { fiveSentenceLoremIpsum } from '../../mock/lipsum'

export default {
  title: 'notification/card/NotificationCard',
  component: NotificationCard,
} as Meta

export const ExampleNotificationCard: Story<NotificationCardProps> = (args) => <NotificationCard {...args} />

ExampleNotificationCard.args = {
  notification: {
    appName: 'Example App',
    timestamp: '2 hours ago',
    title: 'Example Title',
    subtitle: 'Example Subtitle',
    body: fiveSentenceLoremIpsum,
  },
}
