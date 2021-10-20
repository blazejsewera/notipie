import * as React from 'react'
import { Meta, Story } from '@storybook/react'
import type { NotificationCardProps } from '../../../src/component/notification/card/NotificationCard'
import type { Notification } from '../../../src/type/notification'
import { NotificationCard } from '../../../src/component/notification/card/NotificationCard'
import { intlMock } from '../../../src/mock/intl.mock'
import { full, fullWithImage, fullWithLoremIpsum, partial, minimal } from '../../../src/mock/notification.mock'

export default {
  title: 'notification/card/NotificationCard',
  component: NotificationCard,
} as Meta

const NotificationCardStoryFactory = (notification: Notification) => {
  const story: Story<NotificationCardProps> = args => <NotificationCard {...args} />

  story.args = {
    intl: intlMock,
    notification,
  }
  return story
}

const [Full, FullWithImage, FullWithLoremIpsum, Partial, Minimal] = [
  full,
  fullWithImage,
  fullWithLoremIpsum,
  partial,
  minimal,
].map(n => NotificationCardStoryFactory(n))

export { Full, FullWithImage, FullWithLoremIpsum, Partial, Minimal }
