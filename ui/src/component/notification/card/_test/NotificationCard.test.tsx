import { create as render } from 'react-test-renderer'
import { Notification } from '../../../../type/notification'
import { NotificationCard } from '../NotificationCard'
import { intlMock } from '../../../../mock/intl.mock'
import { full, fullWithImage, fullWithLoremIpsum, minimal, partial, read } from '../../../../mock/notification.mock'
import { NotificationCardHandlers } from '../../../../type/handler'

describe('NotificationCard component', () => {
  // given
  const intl = intlMock
  const handlers: NotificationCardHandlers = {
    onArchive: jest.fn(),
    onCheck: jest.fn(),
    onExternal: jest.fn(),
  }

  const testNotificationCard = (notification: Notification) => {
    const tree = render(<NotificationCard {...{ notification, intl, handlers }} />).toJSON()
    expect(tree).toMatchSnapshot()
  }

  it('renders correctly with full notification', () => {
    testNotificationCard(full)
  })
  it('renders correctly with full with image notification', () => {
    testNotificationCard(fullWithImage)
  })
  it('renders correctly with full with very long text notification', () => {
    testNotificationCard(fullWithLoremIpsum)
  })
  it('renders correctly with partial notification', () => {
    testNotificationCard(partial)
  })
  it('renders correctly with minimal notification', () => {
    testNotificationCard(minimal)
  })
  it('renders correctly with read notification', () => {
    testNotificationCard(read)
  })
})
