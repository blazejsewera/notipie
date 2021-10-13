import * as React from 'react'
import type { Notification } from '../../../../type/notification'
import { NotificationCard } from '../NotificationCard'
import renderer from 'react-test-renderer'
import { intlMock } from '../../../../mock/intl.mock'
import { full, fullWithImage, minimal, partial } from '../../../../mock/notification.mock'

describe('NotificationCard component tests', () => {
  // given
  const intl = intlMock

  const testNotificationCard = (notification: Notification) => {
    const tree = renderer.create(<NotificationCard notification={notification} intl={intl} />).toJSON()
    expect(tree).toMatchSnapshot()
  }

  it('renders correctly with minimal notification', () => {
    testNotificationCard(minimal)
  })
  it('renders correctly with partial notification', () => {
    testNotificationCard(partial)
  })
  it('renders correctly with full notification', () => {
    testNotificationCard(full)
  })
  it('renders correctly with full with image notification', () => {
    testNotificationCard(fullWithImage)
  })
})
