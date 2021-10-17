import * as React from 'react'
import { create as render } from 'react-test-renderer'
import type { Notification } from '../../../../type/notification'
import { intlMock } from '../../../../mock/intl.mock'
import { NotificationContainer } from '../NotificationContainer'
import { full, minimal } from '../../../../mock/notification.mock'

describe('NotificationContainer component', () => {
  // given
  const intl = intlMock

  const testNotificationContainer = (notifications: Notification[]) => {
    const tree = render(<NotificationContainer {...{ title: 'TestTitle', notifications, intl }} />).toJSON()
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
