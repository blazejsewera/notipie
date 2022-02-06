import { Meta, Story } from '@storybook/react'
import { NotificationCardProps } from '../../../src/component/notification/card/NotificationCard'
import { NotificationWithHandlers } from '../../../src/type/notification'
import { NotificationCard } from '../../../src/component/notification/card/NotificationCard'
import { intlMock } from '../../../src/mock/intl.mock'
import {
  fullWithHandlers,
  fullWithImageWithHandlers,
  fullWithLoremIpsumWithHandlers,
  partialWithHandlers,
  minimalWithHandlers,
} from '../../../src/mock/notification.mock'

export default {
  title: 'notification/card/NotificationCard',
  component: NotificationCard,
} as Meta

const NotificationCardStoryFactory = (notificationWithHandlers: NotificationWithHandlers) => {
  const story: Story<NotificationCardProps> = args => <NotificationCard {...args} />

  story.args = {
    intl: intlMock,
    notificationWithHandlers,
  }
  return story
}

const [Full, FullWithImage, FullWithLoremIpsum, Partial, Minimal] = [
  fullWithHandlers,
  fullWithImageWithHandlers,
  fullWithLoremIpsumWithHandlers,
  partialWithHandlers,
  minimalWithHandlers,
].map(n => NotificationCardStoryFactory(n))

export { Full, FullWithImage, FullWithLoremIpsum, Partial, Minimal }
