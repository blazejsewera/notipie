import * as React from 'react'
import { Meta, Story } from '@storybook/react'
import { NotificationContainerProps } from '../../../src/component/notification/container/NotificationContainer'
import { NotificationContainer } from '../../../src/component/notification/container/NotificationContainer'
import { intlMock } from '../../../src/mock/intl.mock'
import { full, fullWithLoremIpsum, minimal, partial } from '../../../src/mock/notification.mock'
import { handlers as mockContainerHandlers } from '../../../src/mock/notificationContainer.mock'

export default {
  title: 'notification/container/NotificationContainer',
  component: NotificationContainer,
} as Meta

export const ExampleNotificationContainer: Story<NotificationContainerProps> = args => (
  <NotificationContainer {...args} />
)

ExampleNotificationContainer.args = {
  title: 'Tag title',
  intl: intlMock,
  notifications: [full, fullWithLoremIpsum, partial, minimal],
  handlers: mockContainerHandlers,
}
