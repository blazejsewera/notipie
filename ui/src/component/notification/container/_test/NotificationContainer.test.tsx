import { create as render } from 'react-test-renderer'
import { NotificationWithHandlers } from '../../../../type/notification'
import { intlMock } from '../../../../mock/intl.mock'
import { NotificationContainer } from '../NotificationContainer'
import { fullWithHandlers, minimalWithHandlers } from '../../../../mock/notification.mock'
import { NotificationContainerHandlers } from '../../../../type/handler'

describe('NotificationContainer component', () => {
  // given
  const intl = intlMock
  const handlers: NotificationContainerHandlers = { onCheckAll: jest.fn() }

  const testNotificationContainer = (notificationsWithHandlers: NotificationWithHandlers[]) => {
    const tree = render(
      <NotificationContainer {...{ title: 'TestTitle', notificationsWithHandlers, intl, handlers }} />,
    ).toJSON()
    expect(tree).toMatchSnapshot()
  }

  it('renders correctly with no notifications', () => {
    testNotificationContainer([])
  })
  it('renders correctly with one notification', () => {
    testNotificationContainer([fullWithHandlers])
  })
  it('renders correctly with multiple notifications', () => {
    testNotificationContainer([fullWithHandlers, minimalWithHandlers])
  })
})
