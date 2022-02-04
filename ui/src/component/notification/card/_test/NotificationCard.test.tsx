import { create as render } from 'react-test-renderer'
import { NotificationWithHandlers } from '../../../../type/notification'
import { NotificationCard } from '../NotificationCard'
import { intlMock } from '../../../../mock/intl.mock'
import {
  fullWithHandlers,
  fullWithImageWithHandlers,
  fullWithLoremIpsumWithHandlers,
  minimalWithHandlers,
  partialWithHandlers,
} from '../../../../mock/notification.mock'
import { NotificationCardHandlers } from '../../../../type/handler'

describe('NotificationCard component', () => {
  // given
  const intl = intlMock
  const handlers: NotificationCardHandlers = {
    onArchive: jest.fn(),
    onCheck: jest.fn(),
    onSettings: jest.fn(),
  }

  const testNotificationCard = (notificationWithHandlers: NotificationWithHandlers) => {
    const tree = render(<NotificationCard {...{ notificationWithHandlers, intl, handlers }} />).toJSON()
    expect(tree).toMatchSnapshot()
  }

  it('renders correctly with full notification', () => {
    testNotificationCard(fullWithHandlers)
  })
  it('renders correctly with full with image notification', () => {
    testNotificationCard(fullWithImageWithHandlers)
  })
  it('renders correctly with full with very long text notification', () => {
    testNotificationCard(fullWithLoremIpsumWithHandlers)
  })
  it('renders correctly with partial notification', () => {
    testNotificationCard(partialWithHandlers)
  })
  it('renders correctly with minimal notification', () => {
    testNotificationCard(minimalWithHandlers)
  })
})
