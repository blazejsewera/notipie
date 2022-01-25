import { create as render } from 'react-test-renderer'
import { Notification } from '../../../../type/notification'
import { intlMock } from '../../../../mock/intl.mock'
import { NotificationContainer } from '../NotificationContainer'
import { full, minimal } from '../../../../mock/notification.mock'
import { NotificationContainerHandlers } from '../../../../type/handler'

describe('NotificationContainer component', () => {
  // given
  const intl = intlMock
  const handlers: NotificationContainerHandlers = { onCheckAll: jest.fn() }

  const testNotificationContainer = (notifications: Notification[]) => {
    const tree = render(<NotificationContainer {...{ title: 'TestTitle', notifications, intl, handlers }} />).toJSON()
    expect(tree).toMatchSnapshot()
  }

  it('renders correctly with no notifications', () => {
    testNotificationContainer([])
  })
  it('renders correctly with one notification', () => {
    testNotificationContainer([full])
  })
  it('renders correctly with multiple notifications', () => {
    testNotificationContainer([full, minimal])
  })
})
