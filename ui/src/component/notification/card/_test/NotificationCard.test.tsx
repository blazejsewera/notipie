import * as React from 'react'
import { create as render } from 'react-test-renderer'
import type { Notification } from '../../../../type/notification'
import { NotificationCard } from '../NotificationCard'
import { intlMock } from '../../../../mock/intl.mock'
import { full, fullWithImage, fullWithLoremIpsum, minimal, partial } from '../../../../mock/notification.mock'

describe('NotificationCard component', () => {
  // given
  const intl = intlMock

  const testNotificationCard = (notification: Notification) => {
    const tree = render(<NotificationCard notification={notification} intl={intl} />).toJSON()
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
})
