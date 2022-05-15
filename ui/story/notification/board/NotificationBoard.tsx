import { Meta, Story } from '@storybook/react'
import { NotificationBoard, NotificationBoardProps } from '../../../src/component/notification/board/NotificationBoard'
import { intlMock as intl } from '../../../src/mock/intl.mock'
import { full, fullWithLoremIpsum, minimal, otherApp, partial } from '../../../src/mock/notification.mock'

export default {
  title: 'notification/board/NotificationBoard',
  component: NotificationBoard,
} as Meta

export const ExampleNotificationBoard: Story<NotificationBoardProps> = args => <NotificationBoard {...args} />

ExampleNotificationBoard.args = {
  notifications: [full, fullWithLoremIpsum, partial, minimal, otherApp],
  intl,
  containerHandlersFactory: () => ({ onCheckAll: () => {} }),
}
