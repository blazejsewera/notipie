import * as React from 'react'
import { Meta, Story } from '@storybook/react'
import type { NotificationContainerProps } from '../../../src/component/notification/container/NotificationContainer'
import { NotificationContainer } from '../../../src/component/notification/container/NotificationContainer'
import { fiveSentenceLoremIpsum } from '../../../src/mock/asset/text/lipsum'
import { intlMock } from '../../../src/mock/intl.mock'

export default {
  title: 'notification/container/NotificationContainer',
  component: NotificationContainer,
} as Meta

export const ExampleNotificationContainer: Story<NotificationContainerProps> = (args) => (
  <NotificationContainer {...args} />
)

ExampleNotificationContainer.args = {
  title: 'Tag title',
  intl: intlMock,
  notifications: [
    { appName: 'A1', title: 'Title 1', timestamp: '2 hours ago', id: '1' },
    {
      appName: 'A2',
      title: 'Title 2',
      subtitle: 'Subtitle 2',
      body: fiveSentenceLoremIpsum,
      timestamp: '3 hours ago',
      id: '2',
    },
  ],
}
